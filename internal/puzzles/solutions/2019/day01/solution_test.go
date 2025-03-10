package day01

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/advent-of-code/internal/puzzles/common/utils"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2019"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "1"
	got := s.Day()

	assert.Equal(t, want, got)
}

func Test_module_fuel(t *testing.T) {
	type fields struct {
		mass int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "mass 12",
			fields: fields{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "mass 14",
			fields: fields{
				mass: 14,
			},
			want: 2,
		},
		{
			name: "mass 1969",
			fields: fields{
				mass: 1969,
			},
			want: 654,
		},
		{
			name: "mass 100756",
			fields: fields{
				mass: 100756,
			},
			want: 33583,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := module{
				mass: tt.fields.mass,
			}

			got := m.fuel()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calcPart1(t *testing.T) {
	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calcPart1(in, res, done)

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "mass 12",
			input: 12,
			want:  2,
		},
		{
			name:  "mass 14",
			input: 14,
			want:  2,
		},
		{
			name:  "mass 1969",
			input: 1969,
			want:  654,
		},
		{
			name:  "mass 100756",
			input: 100756,
			want:  33583,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in <- module{mass: tt.input}

			var got int

		LOOP:
			for {
				select {
				case r := <-res:
					got += r
				case <-done:
					break LOOP
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}

	close(in)
}

func Test_calcPart2(t *testing.T) {
	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calcPart2(in, res, done)

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "mass 12",
			input: 12,
			want:  2,
		},
		{
			name:  "mass 14",
			input: 14,
			want:  2,
		},
		{
			name:  "mass 1969",
			input: 1969,
			want:  966,
		},
		{
			name:  "mass 100756",
			input: 100756,
			want:  50346,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in <- module{mass: tt.input}

			var got int

		LOOP:
			for {
				select {
				case r := <-res:
					got += r
				case <-done:
					break LOOP
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}

	close(in)
}

func Test_calc(t *testing.T) {
	type args struct {
		inputPath string
		calcFn    calcFunc
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "part 1",
			args: args{
				inputPath: filepath.Join("testdata", "input.txt"),
				calcFn:    calcPart1,
			},
			want:    "34241",
			wantErr: false,
		},
		{
			name: "part 2",
			args: args{
				inputPath: filepath.Join("testdata", "input.txt"),
				calcFn:    calcPart2,
			},
			want:    "51316",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.ReaderFromFile(t, tt.args.inputPath)
			got, err := calc(input, tt.args.calcFn)

			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
