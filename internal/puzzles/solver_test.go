package puzzles_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

type mockSolver struct {
	year string
	name string
}

func (m mockSolver) Year() string {
	return m.year
}

func (m mockSolver) Part1(_ io.Reader) (string, error) {
	return "part 1 of mockSolver", nil
}

func (m mockSolver) Part2(_ io.Reader) (string, error) {
	return "part 2 of mockSolver", nil
}

func (m mockSolver) Day() string {
	return m.name
}

type anotherMockSolver struct {
	year string
	name string
}

func (a anotherMockSolver) Year() string {
	return a.year
}

func (a anotherMockSolver) Part1(_ io.Reader) (string, error) {
	return "part 1 of anotherMockSolver", nil
}

func (a anotherMockSolver) Part2(_ io.Reader) (string, error) {
	return "part 2 of anotherMockSolver", nil
}

func (a anotherMockSolver) Day() string {
	return a.name
}

func makeAndRegisterSolvers(tb testing.TB) {
	solvers := map[string]map[string]puzzles.Solver{
		"2019": {
			"mock": mockSolver{
				year: "2019",
				name: "mock",
			},
			"anotherMock": anotherMockSolver{
				year: "2019",
				name: "anotherMock",
			},
		},
		"2017": {
			"mock1": mockSolver{
				year: "2017",
				name: "mock1",
			},
		},
	}

	for _, solversYear := range solvers {
		for _, solver := range solversYear {
			puzzles.Register(solver)
		}
	}

	tb.Cleanup(func() {
		puzzles.UnregisterAllSolvers(tb)
	})
}

func TestGetSolver(t *testing.T) {
	makeAndRegisterSolvers(t)

	// get existing solver
	gotSolver, err := puzzles.GetSolver("2019", "mock")
	assert.NoError(t, err)
	assert.IsType(t, mockSolver{}, gotSolver)

	// get existing solver
	gotSolver, err = puzzles.GetSolver("2017", "mock1")
	assert.NoError(t, err)
	assert.IsType(t, mockSolver{}, gotSolver)

	// get existing solver
	gotSolver, err = puzzles.GetSolver("2019", "anotherMock")
	assert.NoError(t, err)
	assert.IsType(t, anotherMockSolver{}, gotSolver)

	// get not existing solver
	gotSolver, err = puzzles.GetSolver("2018", "not-existed")
	assert.Error(t, err)
	assert.IsType(t, nil, gotSolver)

	// get not existing solver
	gotSolver, err = puzzles.GetSolver("2018", "not-existed")
	assert.Error(t, err)
	assert.IsType(t, nil, gotSolver)

	assert.Panics(t, func() {
		puzzles.UnregisterAllSolvers(nil)
	})
}

func TestRun(t *testing.T) {
	puzzles.Register(mockSolver{
		year: "2019",
		name: "mockSolver",
	})
	defer puzzles.UnregisterAllSolvers(t)

	s, err := puzzles.GetSolver("2019", "mockSolver")
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
				Year:  "2019",
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

func TestSolversYears(t *testing.T) {
	makeAndRegisterSolvers(t)

	expectedYears := []string{"2019", "2017"}
	years := puzzles.GetYears()

	assert.ElementsMatch(t, expectedYears, years)
}

func TestSolversByYear(t *testing.T) {
	puzzles.UnregisterAllSolvers(t)
	defer puzzles.UnregisterAllSolvers(t)

	makeAndRegisterSolvers(t)

	solvers := puzzles.DaysByYear("2019")
	expectedSolvers := []string{"mock", "anotherMock"}

	assert.ElementsMatch(t, expectedSolvers, solvers)

	solvers = puzzles.DaysByYear("2017")
	expectedSolvers = []string{"mock1"}

	assert.ElementsMatch(t, expectedSolvers, solvers)
}
