package day06

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Year(t *testing.T) {
	var s solution

	want := "2021"
	got := s.Year()

	assert.Equal(t, want, got)
}

func Test_solution_Day(t *testing.T) {
	var s solution

	want := "6"
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
				input: strings.NewReader("3,4,3,1,2"),
			},
			want:    "5934",
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
				input: strings.NewReader(""),
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

func Test_school_population(t *testing.T) {
	type args struct {
		days         int
		initialState []int
	}

	type expected struct {
		wantFishes []int
	}

	type test struct {
		name string
		args args
		want expected
	}

	tests := []test{
		{
			name: "After  1 day:  2,3,2,0,1",
			args: args{
				days:         1,
				initialState: []int{3, 4, 3, 1, 2},
			},
			want: expected{
				wantFishes: []int{2, 3, 2, 0, 1},
			},
		},
		{
			name: "After  2 days: 1,2,1,6,0,8",
			args: args{
				days:         2,
				initialState: []int{3, 4, 3, 1, 2},
			},
			want: expected{
				wantFishes: []int{1, 2, 1, 6, 0, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sch := newSchool(tt.args.days)

			sch.addElderFishes(tt.args.initialState)

			<-sch.populate()

			assert.ElementsMatch(t, tt.want.wantFishes, sch.getFishes())
		})

	}

}
