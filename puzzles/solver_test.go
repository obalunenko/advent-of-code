package puzzles_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

type mockSolver struct{}

func (m mockSolver) Part1(input io.Reader) (string, error) {
	return fmt.Sprint("part 1 of mockSolver"), nil
}

func (m mockSolver) Part2(input io.Reader) (string, error) {
	return fmt.Sprint("part 2 of mockSolver"), nil
}

func (m mockSolver) Name() string {
	return "mockSolver"
}

type anotherMockSolver struct{}

func (a anotherMockSolver) Part1(input io.Reader) (string, error) {
	return fmt.Sprint("part 1 of anotherMockSolver"), nil
}

func (a anotherMockSolver) Part2(input io.Reader) (string, error) {
	return fmt.Sprint("part 2 of anotherMockSolver"), nil
}

func (a anotherMockSolver) Name() string {
	return "anotherMockSolver"
}

func TestSolver(t *testing.T) {
	defer puzzles.UnregisterAllSolvers(t)

	solvers := map[string]puzzles.Solver{
		"mock":        mockSolver{},
		"anotherMock": anotherMockSolver{},
	}

	solversList := make([]string, 0, len(solvers))

	for name, solver := range solvers {
		puzzles.Register(name, solver)
		solversList = append(solversList, name)
	}

	// get existing solver
	gotSolver, err := puzzles.GetSolver("mock")
	assert.NoError(t, err)
	assert.IsType(t, mockSolver{}, gotSolver)

	// get existing solver
	gotSolver, err = puzzles.GetSolver("anotherMock")
	assert.NoError(t, err)
	assert.IsType(t, anotherMockSolver{}, gotSolver)

	// get not existing solver
	gotSolver, err = puzzles.GetSolver("not-existed")
	assert.Error(t, err)
	assert.IsType(t, nil, gotSolver)

	// get solvers list
	gotList := puzzles.Solvers()
	assert.ElementsMatch(t, solversList, gotList)

	assert.Panics(t, func() {
		puzzles.UnregisterAllSolvers(nil)
	})
}
