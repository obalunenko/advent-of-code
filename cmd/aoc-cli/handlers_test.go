package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func Test_makeMenuItemsList(t *testing.T) {
	type args struct {
		list     []string
		commands []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "without commands",
			args: args{
				list:     []string{"1", "2", "3"},
				commands: nil,
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "with commands",
			args: args{
				list:     []string{"1", "2", "3"},
				commands: []string{"cmd1", "cmd2"},
			},
			want: []string{"1", "2", "3", "cmd1", "cmd2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeMenuItemsList(tt.args.list, tt.args.commands...)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_searcher(t *testing.T) {
	items := makeMenuItemsList([]string{"one", "two", "three"}, exit)

	s := searcher(items)

	assert.True(t, s("o", 0))

	assert.Panics(t, func() {
		s("o", 10)
	})

	assert.True(t, s("t", 2))

	assert.False(t, s("1", 2))
}

func Test_getUrl(t *testing.T) {
	type args struct {
		year string
		day  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				year: puzzles.Year2022.String(),
				day:  puzzles.Day01.String(),
			},
			want: "https://adventofcode.com/2022/day/1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getURL(tt.args.year, tt.args.day)

			assert.Equalf(t, tt.want, got,
				"getURL(%v, %v)", tt.args.year, tt.args.day)
		})
	}
}
