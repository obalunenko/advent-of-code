package day03

import (
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2019"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "3"
	got := s.Day()

	assert.Equal(t, want, got)
}

func Test_solution_Part1(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "distance 6",
			args: args{
				input: strings.NewReader("U7,R6,D4,L4\n" +
					"R8,U5,L5,D3"),
			},
			want:    "6",
			wantErr: false,
		},
		{
			name: "distance 159",
			args: args{
				input: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
					"U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want:    "159",
			wantErr: false,
		},
		{
			name: "distance 159",
			args: args{
				input: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
					"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want:    "135",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findCross(t *testing.T) {
	type args struct {
		wm1 map[pos]int
		wm2 map[pos]int
	}

	tests := []struct {
		name string
		args args
		want []pos
	}{
		{
			name: "",
			args: args{
				wm1: map[pos]int{
					{x: 0, y: 1}: 1,
					{x: 3, y: 1}: 1,
					{x: 4, y: 1}: 1,
					{x: 5, y: 2}: 1,
					{x: 6, y: 3}: 1,
				},
				wm2: map[pos]int{
					{x: 0, y: 1}: 1,
					{x: 4, y: 1}: 1,
					{x: 6, y: 3}: 1,
				},
			},
			want: []pos{
				{x: 0, y: 1},
				{x: 4, y: 1},
				{x: 6, y: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCross(tt.args.wm1, tt.args.wm2)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_solution_Part2(t *testing.T) {
	var s solution

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "610",
			args: args{
				input: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
					"U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want:    "610",
			wantErr: false,
		},
		{
			name: "410",
			args: args{
				input: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
					"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want:    "410",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part2(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
