package day02

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

	want := "2"
	got := s.Day()

	assert.Equal(t, want, got)
}

func Test_nounVerb(t *testing.T) {
	type args struct {
		noun int
		verb int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				noun: 12,
				verb: 2,
			},
			want: 1202,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := nounVerb(tt.args.noun, tt.args.verb)
			assert.Equal(t, tt.want, got)
		})
	}
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
				input: strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50,2,3"),
			},
			want:    "200",
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
