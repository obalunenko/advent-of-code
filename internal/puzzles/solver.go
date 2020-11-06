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

var (
	// ErrNotImplemented signal that puzzle in not implemented yet.
	ErrNotImplemented = errors.New("not implemented")
)

// Solver represents solutions for puzzles methods.
type Solver interface {
	Part1(input io.Reader) (string, error)
	Part2(input io.Reader) (string, error)
	Name() string
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
	name := solver.Name()

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

// NamesByYear returns a sorted list of the names of the registered puzzle solvers for passed year.
func NamesByYear(year string) []string {
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

// GetSolver returns registered solver by passed puzzle name.
func GetSolver(year string, name string) (Solver, error) {
	if year == "" {
		return nil, errors.New("empty puzzle year")
	}

	if name == "" {
		return nil, errors.New("empty puzzle name")
	}

	solversMu.Lock()
	defer solversMu.Unlock()

	solversYaer, exist := solvers[year]

	if !exist {
		return nil, fmt.Errorf("unknown puzzle year [%s]", year)
	}

	s, exist := solversYaer[name]
	if !exist {
		return nil, fmt.Errorf("unknown puzzle name [%s]", name)
	}

	return s, nil
}

// Result represents puzzle solution result.
type Result struct {
	Year  string
	Name  string
	Part1 string
	Part2 string
}

// Run uses solver of puzzle and path to input.
func Run(solver Solver, input io.Reader) (Result, error) {
	var (
		err error
	)

	res := Result{
		Year:  solver.Year(),
		Name:  solver.Name(),
		Part1: "",
		Part2: "",
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(input); err != nil {
		return Result{}, fmt.Errorf("failed to read: %w", err)
	}

	b := buf.Bytes()

	res.Part1, err = solver.Part1(bytes.NewReader(b))
	if err != nil && !errors.Is(err, ErrNotImplemented) {
		return Result{}, fmt.Errorf("failed to solve Part1: %w", err)
	}

	res.Part2, err = solver.Part2(bytes.NewReader(b))
	if err != nil && !errors.Is(err, ErrNotImplemented) {
		return Result{}, fmt.Errorf("failed to solve Part2: %w", err)
	}

	return res, nil
}
