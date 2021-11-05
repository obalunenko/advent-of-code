// Package input provides access to embedded puzzles input files.
package input

import (
	"embed"
	"path/filepath"
)

const (
	dir = "data"
)

// content holds our puzzles inputs content.
//go:embed data/*
var content embed.FS

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	return content.ReadFile(filepath.Clean(
		filepath.Join(dir, name)))
}

// MustAsset loads and returns the asset for the given name.
// It panics if the asset could not be found or
// could not be loaded.
func MustAsset(name string) []byte {
	res, err := Asset(name)
	if err != nil {
		panic(err)
	}

	return res
}
