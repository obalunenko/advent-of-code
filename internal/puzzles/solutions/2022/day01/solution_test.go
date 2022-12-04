package day01

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2022"
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
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test example from description",
			args: args{
				input: strings.NewReader("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"),
			},
			want:    "24000",
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
			name: "",
			args: args{
				input: strings.NewReader("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"),
			},
			want:    "45000",
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

func Test_makeElvesList(t *testing.T) {
	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    elves
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "",
			args: args{
				input: strings.NewReader("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"),
			},
			want: elves{
				{
					food:  []int{1000, 2000, 3000},
					total: 0,
				},
				{
					food:  []int{4000},
					total: 0,
				},
				{
					food:  []int{5000, 6000},
					total: 0,
				},
				{
					food:  []int{7000, 8000, 9000},
					total: 0,
				},
				{
					food:  []int{10000},
					total: 0,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeElvesList(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("makeElvesList(%v)", tt.args.input)) {
				return
			}

			assert.Equal(t, tt.want, got, "makeElvesList(%v)", tt.args.input)
		})
	}
}
