package puzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYear_String(t *testing.T) {
	const yearNotExist Year = 99

	var tests = []struct {
		name string
		i    Year
		want string
	}{
		{
			name: "exist",
			i:    Year2017,
			want: "2017",
		},
		{
			name: "not exist",
			i:    yearNotExist,
			want: "Year(99)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
