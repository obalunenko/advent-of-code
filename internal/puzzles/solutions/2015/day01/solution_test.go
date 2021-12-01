package day01

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

	want := "1"
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
			name: "",
			args: args{
				input: strings.NewReader("(())"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("()()"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("((("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("(()(()("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("))((((("),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("())"),
			},
			want:    "-1",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("))("),
			},
			want:    "-1",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader(")))"),
			},
			want:    "-3",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader(")())())"),
			},
			want:    "-3",
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
			name: "",
			args: args{
				input: strings.NewReader(")"),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input: strings.NewReader("()())"),
			},
			want:    "5",
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
