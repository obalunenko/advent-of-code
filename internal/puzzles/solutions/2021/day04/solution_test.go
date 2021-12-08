package day04

import (
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/advent-of-code/internal/puzzles/common/utils"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2021"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "4"
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
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want:    "4512",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part1(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
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
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want:    "",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Part2(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newBingoGame(t *testing.T) {
	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    *bingo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				input: utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")),
			},
			want: &bingo{
				input: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				boards: []*board{
					{
						numbers: [5][5]int{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
					},
					{
						numbers: [5][5]int{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
					},
					{
						numbers: [5][5]int{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newBingoGame(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("newBingoGame(%v)", tt.args.input)) {
				return
			}
			
			assert.Equalf(t, tt.want, got, "newBingoGame(%v)", tt.args.input)
		})
	}
}
