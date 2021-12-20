package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
