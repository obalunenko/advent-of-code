package day08

import (
	"errors"
	"io"
	"path/filepath"
	"testing"
	"testing/iotest"

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

	want := "8"
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
			want:    "26",
			wantErr: assert.NoError,
		},
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: assert.Error,
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
		{
			name: "",
			args: args{
				input: iotest.ErrReader(errors.New("custom error")),
			},
			want:    "",
			wantErr: assert.Error,
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
