package day04

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			name: "",
			args: args{
				input: strings.NewReader("111000-111222"),
			},
			want:    "46",
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

func Test_findPasswords(t *testing.T) {
	type args struct {
		low  string
		high string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				low:  "111000",
				high: "111222",
			},
			want:    46,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := findPasswords(tt.args.low, tt.args.high, isPasswordPart1)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isIncreasing(t *testing.T) {
	type args struct {
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
				n: 790768,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 123456,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				n: 123450,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := isIncreasing(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_intToSlice(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want [6]int
	}{
		{
			name: "",
			args: args{
				n: 123456,
			},
			want: [6]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := intToSlice(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_hasRepeated(t *testing.T) {
	type args struct {
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
				n: 123456,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 122456,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := hasRepeated(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isPassword(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "meets these criteria (double `11`, never decreases)",
			args: args{
				n: 111111,
			},
			want: true,
		},
		{
			name: "does not meet these criteria (decreasing pair of digits `50`)",
			args: args{
				n: 223450,
			},
			want: false,
		},
		{
			name: "does not meet these criteria (no double).",
			args: args{
				n: 123789,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := isPasswordPart1(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isPasswordPart2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "meets these criteria because the digits never decrease and all repeated digits are " +
				"exactly two digits long.",
			args: args{
				n: 112233,
			},
			want: true,
		},
		{
			name: "no longer meets the criteria (the repeated `44` is part of a larger group of `444`)",
			args: args{
				n: 123444,
			},
			want: false,
		},
		{
			name: "meets the criteria (even though `1` is repeated more than twice, it still contains a double `22`).",
			args: args{
				n: 111122,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := isPasswordPart2(tt.args.n)
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
			name: "`111000-111222` produces `8`",
			args: args{
				input: strings.NewReader("111000-111222"),
			},
			want:    "8",
			wantErr: false,
		},
		{
			name: "`111000-111222\n\n` produces `8`",
			args: args{
				input: strings.NewReader("111000-111222\n\n"),
			},
			want:    "8",
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
