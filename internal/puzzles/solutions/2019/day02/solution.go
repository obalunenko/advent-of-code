// Package day02 contains solution for https://adventofcode.com/2019/day/2 puzzle.
package day02

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/common/intcomputer"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2019.String()
}

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	c, err := intcomputer.New(input)
	if err != nil {
		return "", fmt.Errorf("failed to init computer: %w", err)
	}

	const (
		firstPos  = 12
		secondPos = 2
	)

	c.Input(firstPos, secondPos)

	res, err := c.Execute()
	if err != nil {
		return "", fmt.Errorf("failed to calc: %w", err)
	}

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	c, err := intcomputer.New(input)
	if err != nil {
		return "", fmt.Errorf("failed to init computer: %w", err)
	}

	const expected = 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			c.Reset()

			c.Input(i, j)

			res, err := c.Execute()
			if err != nil {
				return "", fmt.Errorf("failed to calc: %w", err)
			}

			if res == expected {
				return strconv.Itoa(nounVerb(i, j)), nil
			}
		}
	}

	return "", errors.New("can't found non and verb")
}

func nounVerb(noun, verb int) int {
	return 100*noun + verb
}
