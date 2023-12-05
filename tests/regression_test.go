package tests_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/obalunenko/getenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/advent-of-code/internal/command"
	"github.com/obalunenko/advent-of-code/internal/puzzles"
	_ "github.com/obalunenko/advent-of-code/internal/puzzles/solutions" // register puzzles solvers.
)

type args struct {
	year string
	name string
}

type testcase struct {
	name    string
	args    args
	want    puzzles.Result
	wantErr bool
}

const (
	regressionEnabled = "AOC_REGRESSION_ENABLED"
)

// Regression tests for all puzzles. Check that answers still correct.
func TestRun(t *testing.T) {
	if !getenv.EnvOrDefault(regressionEnabled, false) {
		t.Skipf("%s disabled", regressionEnabled)
	}

	session := getenv.EnvOrDefault(puzzles.AOCSession, "")
	if session == "" {
		t.Fatalf("%s not set", puzzles.AOCSession)
	}

	ctx := command.ContextWithSession(context.Background(), session)

	var tests []testcase

	tests = append(tests, invalid()...)
	tests = append(tests, testcases2015(t)...)
	tests = append(tests, testcases2016(t)...)
	tests = append(tests, testcases2017(t)...)
	tests = append(tests, testcases2018(t)...)
	tests = append(tests, testcases2019(t)...)
	tests = append(tests, testcases2020(t)...)
	tests = append(tests, testcases2021(t)...)
	tests = append(tests, testcases2022(t)...)
	tests = append(tests, testcases2023(t)...)

	for i := range tests {
		tt := tests[i]

		t.Run(tt.name, func(t *testing.T) {
			got, err := command.Run(ctx, tt.args.year, tt.args.name)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func invalid() []testcase {
	return []testcase{
		{
			name: "empty year",
			args: args{
				year: "",
				name: puzzles.Day01.String(),
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "empty day",
			args: args{
				year: puzzles.Year2016.String(),
				name: "",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "not exist day",
			args: args{
				year: puzzles.Year2016.String(),
				name: "daynotexist",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "not exist year",
			args: args{
				year: "notexist",
				name: puzzles.Day01.String(),
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
	}
}

func tcName(tb testing.TB, year puzzles.Year, day puzzles.Day) string {
	tb.Helper()

	return fmt.Sprintf("%s/%s", year.String(), day.String())
}
