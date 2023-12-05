package solutions

import (
	"errors"
	"testing"

	"github.com/obalunenko/getenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_createNewFromTemplate(t *testing.T) {
	const envName = "AOC_PUZZLE_URL"

	purl, err := getenv.Env[string](envName)
	if err != nil {
		if errors.Is(err, getenv.ErrNotSet) {
			t.Skipf("%s is not set", envName)
		}

		t.Fatalf("failed to get environment variable[%s]: %v", envName, err)
	}

	require.NoError(t, createNewFromTemplate(purl))
}

func Test_parsePuzzleURL(t *testing.T) {
	type args struct {
		url string
	}

	tests := []struct {
		name     string
		args     args
		wandDate puzzleDate
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "valid url",
			args: args{
				url: "https://adventofcode.com/2022/day/1",
			},
			wandDate: puzzleDate{
				year: 2022,
				day:  1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "invalid url",
			args: args{
				url: "https://adventofcode.com/2022",
			},
			wandDate: puzzleDate{
				year: 0,
				day:  0,
			},
			wantErr: assert.Error,
		},
		{
			name: "empty url",
			args: args{
				url: "",
			},
			wandDate: puzzleDate{
				year: 0,
				day:  0,
			},
			wantErr: assert.Error,
		},
		{
			name: "whitespace url",
			args: args{
				url: " ",
			},
			wandDate: puzzleDate{
				year: 0,
				day:  0,
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, err := parsePuzzleURL(tt.args.url)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.wandDate, gotDate)
		})
	}
}
