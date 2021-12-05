package utils

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// ReaderFromFile reads file from fpath and returns content as io.Reader.
// File descriptor will be closed on tests teardown.
func ReaderFromFile(tb testing.TB, fpath string) io.Reader {
	tb.Helper()

	file, err := os.Open(fpath)
	require.NoError(tb, err)

	tb.Cleanup(func() {
		require.NoError(tb, file.Close())
	})

	return file
}
