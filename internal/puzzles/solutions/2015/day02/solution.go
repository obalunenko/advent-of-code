// Package day02 contains solution for https://adventofcode.com/2015/day/2 puzzle.
package day02

import (
	"io"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Year() string {
	return puzzles.Year2015.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
