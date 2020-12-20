package day03

import (
	"io"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day03"
	year       = "2020"
)

type solution struct {
	year string
	name string
}

func (s solution) Year() string {
	return s.year
}

func (s solution) Name() string {
	return s.name
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

func (s solution) Part1(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
