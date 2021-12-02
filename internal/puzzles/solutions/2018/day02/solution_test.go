package day02

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Day(t *testing.T) {
	var s solution

	got := s.Day()
	expected := "2"

	assert.Equal(t, expected, got)
}

func Test_solution_Year(t *testing.T) {
	var s solution

	got := s.Year()
	expected := "2018"

	assert.Equal(t, expected, got)
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
			name: "example from spec",
			args: args{
				input: strings.NewReader("abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab\n"),
			},
			want:    "12",
			wantErr: false,
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
			name: "example from spec",
			args: args{
				input: strings.NewReader("abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz\n"),
			},
			want:    "fgij",
			wantErr: false,
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

func Test_hasNSameLetters(t *testing.T) {
	type args struct {
		s string
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "abcdef",
				n: 2,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "bababc",
				n: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "abcccd",
				n: 3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "abcccd",
				n: 2,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasNSameLetters(tt.args.s, tt.args.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_hasNDiffLetters(t *testing.T) {
	type args struct {
		box1 string
		box2 string
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match",
			args: args{
				box1: "fghij",
				box2: "fguij",
				n:    1,
			},
			want: true,
		},
		{
			name: "not match",
			args: args{
				box1: "abcde",
				box2: "axcye",
				n:    1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasNDiffLetters(tt.args.box1, tt.args.box2, tt.args.n)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getCommonBoxesPart(t *testing.T) {
	type args struct {
		box1 string
		box2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				box1: "abcde",
				box2: "axcye",
			},
			want: "ace",
		},
		{
			name: "",
			args: args{
				box1: "abcde",
				box2: "zxylk",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCommonBoxesPart(tt.args.box1, tt.args.box2)

			assert.Equal(t, tt.want, got)
		})
	}
}
