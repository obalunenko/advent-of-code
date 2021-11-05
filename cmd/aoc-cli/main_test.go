package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
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
func Test_run(t *testing.T) {
	var tests []testcase

	tests = append(tests, invalid()...)
	tests = append(tests, testcases2015()...)
	tests = append(tests, testcases2016()...)
	tests = append(tests, testcases2017()...)
	tests = append(tests, testcases2018()...)
	tests = append(tests, testcases2019()...)
	tests = append(tests, testcases2020()...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args.year, tt.args.name)
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
				name: "day01",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "empty day",
			args: args{
				year: "2016",
				name: "",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "not exist day",
			args: args{
				year: "2016",
				name: "daynotexist",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
		{
			name: "not exist year",
			args: args{
				year: "notexist",
				name: "day01",
			},
			want:    puzzles.Result{},
			wantErr: true,
		},
	}
}

func testcases2015() []testcase {
	const year = "2015"

	return []testcase{
		{
			name: "2015/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "232",
				Part2: "1783",
			},
			wantErr: false,
		},
	}
}

func testcases2016() []testcase {
	const year = "2016"

	return []testcase{
		{
			name: "2016/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "307",
				Part2: "165",
			},
			wantErr: false,
		},
	}
}

func testcases2017() []testcase {
	const year = "2017"

	return []testcase{
		{
			name: "2017/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "1029",
				Part2: "1220",
			},
			wantErr: false,
		},
	}
}

func testcases2018() []testcase {
	const year = "2018"

	return []testcase{
		{
			name: "2018/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "439",
				Part2: "124645",
			},
			wantErr: false,
		},
	}
}

func testcases2019() []testcase {
	const year = "2019"

	return []testcase{
		{
			name: "2019/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "3464458",
				Part2: "5193796",
			},
			wantErr: false,
		},
		{
			name: "2019/day02",
			args: args{
				year: year,
				name: "day02",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day02",
				Part1: "2890696",
				Part2: "8226",
			},
			wantErr: false,
		},
		{
			name: "2019/day03",
			args: args{
				year: year,
				name: "day03",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day03",
				Part1: "1195",
				Part2: "91518",
			},
			wantErr: false,
		},
		{
			name: "2019/day04",
			args: args{
				year: year,
				name: "day04",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day04",
				Part1: "2779",
				Part2: "1972",
			},
			wantErr: false,
		},
	}
}

func testcases2020() []testcase {
	const year = "2020"

	return []testcase{
		{
			name: "2020/day01",
			args: args{
				year: year,
				name: "day01",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day01",
				Part1: "270144",
				Part2: "261342720",
			},
			wantErr: false,
		},
		{
			name: "2020/day02",
			args: args{
				year: year,
				name: "day02",
			},
			want: puzzles.Result{
				Year:  year,
				Name:  "day02",
				Part1: "456",
				Part2: "308",
			},
			wantErr: false,
		},
	}
}
