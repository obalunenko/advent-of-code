package utils

import (
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReaderFromFile(t *testing.T) {
	type args struct {
		tb    testing.TB
		fpath string
	}

	tests := []struct {
		name string
		args args
		want io.Reader
	}{
		{
			name: "",
			args: args{
				tb:    t,
				fpath: filepath.Join("testdata", "reader.txt"),
			},
			want: strings.NewReader("Hello there!\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReaderFromFile(tt.args.tb, tt.args.fpath)

			gotR, err := io.ReadAll(got)
			require.NoError(t, err)

			wantR, err := io.ReadAll(tt.want)
			require.NoError(t, err)

			assert.Equal(t, string(wantR), string(gotR))
		})
	}
}
