package tests_test

import (
	"testing"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func testcases2018(tb testing.TB) []testcase {
	year := puzzles.Year2018

	return []testcase{
		{
			name: tcName(tb, year, puzzles.Day01),
			args: args{
				year: year.String(),
				name: puzzles.Day01.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day01.String(),
				Part1: "439",
				Part2: "124645",
			},
			wantErr: false,
		},
		{
			name: tcName(tb, year, puzzles.Day02),
			args: args{
				year: year.String(),
				name: puzzles.Day02.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day02.String(),
				Part1: "6944",
				Part2: "srijafjzloguvlntqmphenbkd",
			},
			wantErr: false,
		},
		{
			name: tcName(tb, year, puzzles.Day03),
			args: args{
				year: year.String(),
				name: puzzles.Day03.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day03.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day04),
			args: args{
				year: year.String(),
				name: puzzles.Day04.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day04.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day05),
			args: args{
				year: year.String(),
				name: puzzles.Day05.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day05.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day06),
			args: args{
				year: year.String(),
				name: puzzles.Day06.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day06.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day07),
			args: args{
				year: year.String(),
				name: puzzles.Day07.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day07.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day08),
			args: args{
				year: year.String(),
				name: puzzles.Day08.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day08.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day09),
			args: args{
				year: year.String(),
				name: puzzles.Day09.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day09.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day10),
			args: args{
				year: year.String(),
				name: puzzles.Day10.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day10.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day11),
			args: args{
				year: year.String(),
				name: puzzles.Day11.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day11.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day12),
			args: args{
				year: year.String(),
				name: puzzles.Day12.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day12.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day13),
			args: args{
				year: year.String(),
				name: puzzles.Day13.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day13.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day14),
			args: args{
				year: year.String(),
				name: puzzles.Day14.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day14.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day15),
			args: args{
				year: year.String(),
				name: puzzles.Day15.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day15.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day16),
			args: args{
				year: year.String(),
				name: puzzles.Day16.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day16.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day17),
			args: args{
				year: year.String(),
				name: puzzles.Day17.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day17.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day18),
			args: args{
				year: year.String(),
				name: puzzles.Day18.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day18.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day19),
			args: args{
				year: year.String(),
				name: puzzles.Day19.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day19.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day20),
			args: args{
				year: year.String(),
				name: puzzles.Day20.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day20.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day21),
			args: args{
				year: year.String(),
				name: puzzles.Day21.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day21.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day22),
			args: args{
				year: year.String(),
				name: puzzles.Day22.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day22.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day23),
			args: args{
				year: year.String(),
				name: puzzles.Day23.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day23.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day24),
			args: args{
				year: year.String(),
				name: puzzles.Day24.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day24.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
		{
			name: tcName(tb, year, puzzles.Day25),
			args: args{
				year: year.String(),
				name: puzzles.Day25.String(),
			},
			want: puzzles.Result{
				Year:  year.String(),
				Name:  puzzles.Day25.String(),
				Part1: "",
				Part2: "",
			},
			wantErr: true,
		},
	}
}
