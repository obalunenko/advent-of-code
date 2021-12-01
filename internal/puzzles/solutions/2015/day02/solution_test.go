package day02

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2015"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "2"
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
			name: "`2x3x4` requires a total of 58 square feet",
			args: args{
				input: strings.NewReader("2x3x4\n"),
			},
			want:    "58",
			wantErr: false,
		},
		{
			name: "`1x1x10` requires a total of 43 square feet",
			args: args{
				input: strings.NewReader("1x1x10\n"),
			},
			want:    "43",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

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
			name: "`2x3x4` requires a total of 34 square feet",
			args: args{
				input: strings.NewReader("2x3x4\n"),
			},
			want:    "34",
			wantErr: false,
		},
		{
			name: "`1x1x10` requires a total of 14 square feet",
			args: args{
				input: strings.NewReader("1x1x10\n"),
			},
			want:    "14",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

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
