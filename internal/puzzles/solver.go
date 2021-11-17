// Package puzzles declares common interface for puzzle solutions and functionality for register and run them.
package puzzles

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"sync"
	"testing"
)

// Solver represents solutions for puzzles methods.
type Solver interface {
	Part1(input io.Reader) (string, error)
	Part2(input io.Reader) (string, error)
	Day() string
	Year() string
}

var (
	solversMu sync.RWMutex
	solvers   = make(map[string]map[string]Solver)
)

// Register makes a puzzle solver available by the provided name.
// If Register is called twice with the same name or if solver is nil,
// it panics.
func Register(solver Solver) {
	year := solver.Year()
	name := solver.Day()

	solversMu.Lock()
	defer solversMu.Unlock()

	if solver == nil {
		panic("puzzle: Register solver is nil")
	}

	yearSolvers, exist := solvers[year]
	if !exist {
		solvers[year] = make(map[string]Solver)
		yearSolvers = solvers[year]
	}

	if _, dup := yearSolvers[name]; dup {
		panic(fmt.Errorf("puzzle: Register called twice for solver [%s:%s]", year, name))
	}

	yearSolvers[name] = solver
	solvers[year] = yearSolvers
}

// UnregisterAllSolvers cleans up registered solvers. Use for testing only.
func UnregisterAllSolvers(tb testing.TB) {
	if tb == nil {
		panic("UnregisterAllSolvers should be called only inside tests")
	}

	solversMu.Lock()
	defer solversMu.Unlock()

	solvers = make(map[string]map[string]Solver)
}

// DaysByYear returns a sorted list of the days of the registered puzzle solvers for passed year.
func DaysByYear(year string) []string {
	solversMu.RLock()
	defer solversMu.RUnlock()

	list := make([]string, 0, len(solvers[year]))

	for name := range solvers[year] {
		list = append(list, name)
	}

	sort.Strings(list)

	return list
}

// GetYears returns list of available years for solvers.
func GetYears() []string {
	solversMu.RLock()
	defer solversMu.RUnlock()

	list := make([]string, 0, len(solvers))

	for year := range solvers {
		list = append(list, year)
	}

	sort.Strings(list)

	return list
}

// GetSolver returns registered solver by passed puzzle day.
func GetSolver(year, day string) (Solver, error) {
	if year == "" {
		return nil, errors.New("empty puzzle year")
	}

	if day == "" {
		return nil, errors.New("empty puzzle day")
	}

	solversMu.Lock()
	defer solversMu.Unlock()

	solversYear, exist := solvers[year]

	if !exist {
		return nil, fmt.Errorf("unknown puzzle year [%s]", year)
	}

	s, exist := solversYear[day]
	if !exist {
		return nil, fmt.Errorf("unknown puzzle day [%s]", day)
	}

	return s, nil
}

type runParams struct {
	withMetrics metricsFlag
}

// RunOption provides run options pattern.
type RunOption interface {
	Apply(opts *runParams)
}

// WithElapsed add elapsed metric to run options.
func WithElapsed() RunOption {
	return withElapsed{}
}

// WithBenchmark add benchmark metric to run options.
func WithBenchmark() RunOption {
	return withBenchmark{}
}

type withElapsed struct{}

func (w withElapsed) Apply(opts *runParams) {
	opts.withMetrics.AddFlag(metricsFlagElapsed)
}

type withBenchmark struct{}

func (w withBenchmark) Apply(opts *runParams) {
	opts.withMetrics.AddFlag(metricsFlagBenchmark)
}

func makeRunParams(opts []RunOption) runParams {
	var p runParams

	for _, opt := range opts {
		opt.Apply(&p)
	}

	return p
}

// Run uses solver of puzzle and path to input.
func Run(solver Solver, input io.Reader, opts ...RunOption) (Result, error) {
	params := makeRunParams(opts)

	res := Result{
		Year:    solver.Year(),
		Name:    solver.Day(),
		Part1:   unsolved,
		Part2:   unsolved,
		metrics: nil,
	}

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(input); err != nil {
		return Result{}, fmt.Errorf("failed to read: %w", err)
	}

	b := buf.Bytes()

	apply := res.addMetrics(solver, b, params.withMetrics)
	defer apply()

	if err := res.addAnswers(solver, b); err != nil {
		return Result{}, fmt.Errorf("failed to add answers: %w", err)
	}

	return res, nil
}

type applyMetricFunc func()

func (r *Result) addMetrics(solver Solver, input []byte, mf metricsFlag) func() {
	if mf.HasFlag(metricsFlagNone) {
		return func() {
			r.metrics = nil
		}
	}

	const metricsnum = 2

	metricFuncs := make([]applyMetricFunc, 0, metricsnum)

	if mf.HasFlag(metricsFlagElapsed) {
		var em metric

		r.metrics = append(r.metrics, &em)

		metricFuncs = append(metricFuncs, em.elapsed())
	}

	if mf.HasFlag(metricsFlagBenchmark) {
		var bm metric

		r.metrics = append(r.metrics, &bm)

		bf := benchFunc(func() error {
			return r.addAnswers(solver, input)
		})

		metricFuncs = append(metricFuncs, bm.bench(bf))
	}

	return func() {
		for _, f := range metricFuncs {
			f()
		}
	}
}

func (r *Result) addAnswers(s Solver, input []byte) error {
	part1, err := s.Part1(bytes.NewReader(input))
	if err != nil && !errors.Is(err, ErrNotImplemented) {
		return fmt.Errorf("failed to solve Part1: %w", err)
	}

	part2, err := s.Part2(bytes.NewReader(input))
	if err != nil && !errors.Is(err, ErrNotImplemented) {
		return fmt.Errorf("failed to solve Part2: %w", err)
	}

	r.Part1 = part1
	r.Part2 = part2

	return nil
}
