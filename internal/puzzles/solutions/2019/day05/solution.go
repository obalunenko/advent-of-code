package day05

import (
	"io"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day05"
	year       = "2019"
)

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

type solution struct {
	name string
	year string
}

func (s solution) Year() string {
	return s.year
}

func (s solution) Part1(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}
