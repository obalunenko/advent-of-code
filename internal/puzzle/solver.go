package puzzle

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"sync"

	"github.com/pkg/errors"
)

// Solver represents solutions for puzzles methods.
type Solver interface {
	Part1(input io.Reader) (string, error)
	Part2(input io.Reader) (string, error)
	Name() string
}

var (
	solversMu sync.RWMutex
	solvers   = make(map[string]Solver)
)

// Register makes a puzzle solver available by the provided name.
// If Register is called twice with the same name or if solver is nil,
// it panics.
func Register(name string, solver Solver) {
	solversMu.Lock()
	defer solversMu.Unlock()

	if solver == nil {
		panic("puzzle: Register solver is nil")
	}

	if _, dup := solvers[name]; dup {
		panic("puzzle: Register called twice for solver " + name)
	}

	solvers[name] = solver
}

// Solvers returns a sorted list of the names of the registered puzzle solvers.
func Solvers() []string {
	solversMu.RLock()
	defer solversMu.RUnlock()

	list := make([]string, 0, len(solvers))

	for name := range solvers {
		list = append(list, name)
	}

	sort.Strings(list)

	return list
}

// GetSolver returns registered solver by passed puzzle name.
func GetSolver(name string) (Solver, error) {
	if name == "" {
		return nil, errors.New("empty puzzle name")
	}

	solversMu.Lock()
	defer solversMu.Unlock()

	s, exist := solvers[name]

	if !exist {
		return nil, errors.Errorf("unknown puzzle name [%s]", name)
	}

	return s, nil
}

// Run uses solver of puzzle and path to input.
func Run(solver Solver, filepath string) error {
	var (
		input []byte
		res   string
		err   error
	)

	input, err = ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	log.Printf("run puzzle solver [%s]\n", solver.Name())

	res, err = solver.Part1(bytes.NewBuffer(input))
	if err != nil {
		return errors.Wrapf(err, "failed to run Part1 for puzzle [%s]", solver.Name())
	}

	log.Printf("Part1 answer: %s \n", res)

	res, err = solver.Part2(bytes.NewBuffer(input))
	if err != nil {
		return errors.Wrapf(err, "failed to run Part2 for puzzle [%s]", solver.Name())
	}

	log.Printf("Part2 answer: %s \n", res)

	return nil
}
