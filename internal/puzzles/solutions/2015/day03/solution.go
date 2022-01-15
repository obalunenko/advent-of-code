// Package day03 contains solution for https://adventofcode.com/2015/day/3 puzzle.
package day03

import (
	"io"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (solution) Day() string {
	return puzzles.Day03.String()
}

func (solution) Year() string {
	return puzzles.Year2015.String()
}

func (solution) Part1(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
