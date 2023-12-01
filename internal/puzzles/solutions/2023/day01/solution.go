// Package day01 contains solution for https://adventofcode.com/2023/day/1 puzzle.
package day01

import (
	"io"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2023.String()
}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Part1(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
