package day04

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
			want:    "1924",
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
						id: 1,
						numbers: [boardSize][boardSize]number{
							{number{val: 22}, number{val: 13}, number{val: 17}, number{val: 11}, number{val: 0}},
							{number{val: 8}, number{val: 2}, number{val: 23}, number{val: 4}, number{val: 24}},
							{number{val: 21}, number{val: 9}, number{val: 14}, number{val: 16}, number{val: 7}},
							{number{val: 6}, number{val: 10}, number{val: 3}, number{val: 18}, number{val: 5}},
							{number{val: 1}, number{val: 12}, number{val: 20}, number{val: 15}, number{val: 19}},
						},
						state: state{
							verticals:   [boardSize]int{},
							horizontals: [boardSize]int{},
						},
					},
					{
						id: 2,
						numbers: [boardSize][boardSize]number{
							{number{val: 3}, number{val: 15}, number{val: 0}, number{val: 2}, number{val: 22}},
							{number{val: 9}, number{val: 18}, number{val: 13}, number{val: 17}, number{val: 5}},
							{number{val: 19}, number{val: 8}, number{val: 7}, number{val: 25}, number{val: 23}},
							{number{val: 20}, number{val: 11}, number{val: 10}, number{val: 24}, number{val: 4}},
							{number{val: 14}, number{val: 21}, number{val: 16}, number{val: 12}, number{val: 6}},
						},
						state: state{
							verticals:   [boardSize]int{},
							horizontals: [boardSize]int{},
						},
					},
					{
						id: 3,
						numbers: [boardSize][boardSize]number{
							{number{val: 14}, number{val: 21}, number{val: 17}, number{val: 24}, number{val: 4}},
							{number{val: 10}, number{val: 16}, number{val: 15}, number{val: 9}, number{val: 19}},
							{number{val: 18}, number{val: 8}, number{val: 23}, number{val: 26}, number{val: 20}},
							{number{val: 22}, number{val: 11}, number{val: 13}, number{val: 6}, number{val: 5}},
							{number{val: 2}, number{val: 0}, number{val: 12}, number{val: 3}, number{val: 7}},
						},
						state: state{
							verticals:   [boardSize]int{},
							horizontals: [boardSize]int{},
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

func Test_bingo_start(t *testing.T) {
	ctx := context.Background()

	bgame, err := newBingoGame(utils.ReaderFromFile(t, filepath.Join("testdata", "input.txt")))
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		wr  winRule
	}

	type expected struct {
		board *board
		num   int
	}

	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "",
			args: args{
				ctx: ctx,
				wr:  rule(1),
			},
			expected: expected{
				board: &board{
					id: 3,
					numbers: [5][5]number{
						{
							number{val: 14, isMarked: true},
							number{val: 21, isMarked: true},
							number{val: 17, isMarked: true},
							number{val: 24, isMarked: true},
							number{val: 4, isMarked: true},
						},
						{
							number{val: 10},
							number{val: 16},
							number{val: 15},
							number{val: 9, isMarked: true},
							number{val: 19},
						},
						{
							number{val: 18},
							number{val: 8},
							number{val: 23, isMarked: true},
							number{val: 26},
							number{val: 20},
						},
						{
							number{val: 22},
							number{val: 11, isMarked: true},
							number{val: 13},
							number{val: 6},
							number{val: 5, isMarked: true},
						},
						{
							number{val: 2, isMarked: true},
							number{val: 0, isMarked: true},
							number{val: 12},
							number{val: 3},
							number{val: 7, isMarked: true},
						},
					},
					state: state{
						verticals: [boardSize]int{
							0: 2,
							1: 3,
							2: 2,
							3: 2,
							4: 3,
						},
						horizontals: [boardSize]int{
							0: 5,
							1: 1,
							2: 1,
							3: 2,
							4: 3,
						},
					},
				},
				num: 24,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bgame

			gotBoard, gotNum := b.start(tt.args.ctx, tt.args.wr)

			equalBoards(t, tt.expected.board, gotBoard)
			assert.Equal(t, tt.expected.num, gotNum)
		})
	}
}

func equalBoards(t testing.TB, expected, got *board) {
	assert.Equal(t, expected.numbers, got.numbers, "numbers")

	assert.Equal(t, expected.id, got.id, "id")

	assert.Equal(t, fmt.Sprint(expected.state.horizontals), fmt.Sprint(got.state.horizontals), "horizontals")

	assert.Equal(t, fmt.Sprint(expected.state.verticals), fmt.Sprint(got.state.verticals), "verticals")
}

func Test_board_sumMarked(t *testing.T) {
	b := board{
		id: 0,
		numbers: [5][5]number{
			{
				number{val: 14, isMarked: true},
				number{val: 21, isMarked: true},
				number{val: 17, isMarked: true},
				number{val: 24, isMarked: true},
				number{val: 4, isMarked: true},
			},
			{
				number{val: 10},
				number{val: 16},
				number{val: 15},
				number{val: 9, isMarked: true},
				number{val: 19},
			},
			{
				number{val: 18},
				number{val: 8},
				number{val: 23, isMarked: true},
				number{val: 26},
				number{val: 20},
			},
			{
				number{val: 22},
				number{val: 11, isMarked: true},
				number{val: 13},
				number{val: 6},
				number{val: 5, isMarked: true},
			},
			{
				number{val: 2, isMarked: true},
				number{val: 0, isMarked: true},
				number{val: 12},
				number{val: 3},
				number{val: 7, isMarked: true},
			},
		},
		state: state{
			verticals: [boardSize]int{
				0: 2,
				1: 3,
				2: 2,
				3: 2,
				4: 3,
			},
			horizontals: [boardSize]int{
				0: 5,
				1: 1,
				2: 1,
				3: 2,
				4: 3,
			},
		},
	}

	tests := []struct {
		name  string
		board board
		want  int
	}{
		{
			name:  "",
			board: b,
			want:  188,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.board.sumMarked(), "sumMarked()")
		})
	}
}
