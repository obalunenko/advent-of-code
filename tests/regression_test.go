package tests_test

import (
	"context"
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

// Regression tests for all puzzles. Check that answers still correct.
func TestRun(t *testing.T) {
	if !getenv.BoolOrDefault("AOC_REGRESSION_ENABLED", false) {
		t.Skip("Regression test disabled")
	}

	session := getenv.StringOrDefault("AOC_SESSION", "")
	if session == "" {
		t.Fatal("AOC_SESSION not set")
	}

	ctx := command.ContextWithSession(context.Background(), session)

	var tests []testcase

	tests = append(tests, invalid()...)
	tests = append(tests, testcases2015()...)
	tests = append(tests, testcases2016()...)
	tests = append(tests, testcases2017()...)
	tests = append(tests, testcases2018()...)
	tests = append(tests, testcases2019()...)
	tests = append(tests, testcases2020()...)
	tests = append(tests, testcases2021()...)

	for _, tt := range tests {
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

func testcases2015() []testcase {
	year := puzzles.Year2015.String()

	return []testcase{
		{
			name: "2015/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "232",
				Part2: "1783",
			},
			wantErr: false,
		},
		{
			name: "2015/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "1598415",
				Part2: "3812909",
			},
			wantErr: false,
		},
		{
			name: "2015/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2015/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2016() []testcase {
	year := puzzles.Year2016.String()

	return []testcase{
		{
			name: "2016/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "307",
				Part2: "165",
			},
			wantErr: false,
		},
		{
			name: "2016/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "48584",
				Part2: "563B6",
			},
			wantErr: false,
		},
		{
			name: "2016/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2016/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2017() []testcase {
	year := puzzles.Year2017.String()

	return []testcase{
		{
			name: "2017/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "1029",
				Part2: "1220",
			},
			wantErr: false,
		},
		{
			name: "2017/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "32020",
				Part2: "236",
			},
			wantErr: false,
		},
		{
			name: "2017/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2017/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2018() []testcase {
	year := puzzles.Year2018.String()

	return []testcase{
		{
			name: "2018/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "439",
				Part2: "124645",
			},
			wantErr: false,
		},
		{
			name: "2018/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "6944",
				Part2: "srijafjzloguvlntqmphenbkd",
			},
			wantErr: false,
		},
		{
			name: "2018/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2018/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2019() []testcase {
	year := puzzles.Year2019.String()

	return []testcase{
		{
			name: "2019/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "3464458",
				Part2: "5193796",
			},
			wantErr: false,
		},
		{
			name: "2019/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "2890696",
				Part2: "8226",
			},
			wantErr: false,
		},
		{
			name: "2019/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "1195",
				Part2: "91518",
			},
			wantErr: false,
		},
		{
			name: "2019/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "2779",
				Part2: "1972",
			},
			wantErr: false,
		},
		{
			name: "2019/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2019/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2020() []testcase {
	year := puzzles.Year2020.String()

	return []testcase{
		{
			name: "2020/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "270144",
				Part2: "261342720",
			},
			wantErr: false,
		},
		{
			name: "2020/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "456",
				Part2: "308",
			},
			wantErr: false,
		},
		{
			name: "2020/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2020/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}

func testcases2021() []testcase {
	year := puzzles.Year2021.String()

	return []testcase{
		{
			name: "2021/day01",
			args: args{
				year: year,
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day01.String(),
				Part1: "1482",
				Part2: "1518",
			},
			wantErr: false,
		},
		{
			name: "2021/day02",
			args: args{
				year: year,
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day02.String(),
				Part1: "1484118",
				Part2: "1463827010",
			},
			wantErr: false,
		},
		{
			name: "2021/day03",
			args: args{
				year: year,
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: false,
		},
		{
			name: "2021/day04",
			args: args{
				year: year,
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day05",
			args: args{
				year: year,
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day06",
			args: args{
				year: year,
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day07",
			args: args{
				year: year,
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day08",
			args: args{
				year: year,
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day09",
			args: args{
				year: year,
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day10",
			args: args{
				year: year,
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day11",
			args: args{
				year: year,
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day12",
			args: args{
				year: year,
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day13",
			args: args{
				year: year,
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day14",
			args: args{
				year: year,
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day15",
			args: args{
				year: year,
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day16",
			args: args{
				year: year,
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day17",
			args: args{
				year: year,
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day18",
			args: args{
				year: year,
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day19",
			args: args{
				year: year,
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day20",
			args: args{
				year: year,
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day21",
			args: args{
				year: year,
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day22",
			args: args{
				year: year,
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day23",
			args: args{
				year: year,
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day24",
			args: args{
				year: year,
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: "2021/day25",
			args: args{
				year: year,
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year,
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}
