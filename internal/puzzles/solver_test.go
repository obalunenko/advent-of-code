package puzzles_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oleg-balunenko/advent-of-code/internal/puzzles"
)

type mockSolver struct{}

func (m mockSolver) Part1(_ io.Reader) (string, error) {
	return "part 1 of mockSolver", nil
}

func (m mockSolver) Part2(_ io.Reader) (string, error) {
	return "part 2 of mockSolver", nil
}

func (m mockSolver) Name() string {
	return "mockSolver"
}

type anotherMockSolver struct{}

func (a anotherMockSolver) Part1(_ io.Reader) (string, error) {
	return "part 1 of anotherMockSolver", nil
}

func (a anotherMockSolver) Part2(_ io.Reader) (string, error) {
	return "part 2 of anotherMockSolver", nil
}

func (a anotherMockSolver) Name() string {
	return "anotherMockSolver"
}

func TestSolver(t *testing.T) {
	puzzles.UnregisterAllSolvers(t)
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

func TestRun(t *testing.T) {
	puzzles.Register("mockSolver", mockSolver{})
	defer puzzles.UnregisterAllSolvers(t)

	s, err := puzzles.GetSolver("mockSolver")
	require.NoError(t, err)

	type args struct {
		solver puzzles.Solver
		input  io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    puzzles.Result
		wantErr bool
	}{
		{
			name: "",
			args: args{
				solver: s,
				input:  strings.NewReader("testdata"),
			},
			want: puzzles.Result{
				Name:  "mockSolver",
				Part1: "part 1 of mockSolver",
				Part2: "part 2 of mockSolver",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := puzzles.Run(tt.args.solver, tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
