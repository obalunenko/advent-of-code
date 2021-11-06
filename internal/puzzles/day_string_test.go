package puzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay_String(t *testing.T) {
	const dayNotExist Day = 99

	var tests = []struct {
		name string
		i    Day
		want string
	}{
		{
			name: "exist",
			i:    Day01,
			want: "day01",
		},
		{
			name: "not exist",
			i:    dayNotExist,
			want: "Day(99)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
