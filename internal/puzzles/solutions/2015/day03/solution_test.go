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

	want := "2015"
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
			name: "> delivers presents to 2 houses: one at the starting location, and one to the east.",
			args: args{
				input: strings.NewReader(">\n"),
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.",
			args: args{
				input: strings.NewReader("^>v<\n"),
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: "^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.",
			args: args{
				input: strings.NewReader("^v^v^v^v^v\n"),
			},
			want:    "2",
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
			name: "^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.",
			args: args{
				input: strings.NewReader("^v\n"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.",
			args: args{
				input: strings.NewReader("^>v<\n"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.",
			args: args{
				input: strings.NewReader("^v^v^v^v^v\n"),
			},
			want:    "11",
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
