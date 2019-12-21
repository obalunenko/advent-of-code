package day04

import (
	"io"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

func init() {
	puzzleName, err := puzzles.MakeName("2019", "day04")
	if err != nil {
		panic(err)
	}

	puzzles.Register(puzzleName, solution{
		name: puzzleName,
	})
}

type solution struct {
	name string
}

func (s solution) Part1(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}
