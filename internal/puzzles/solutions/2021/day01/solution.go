// Package day01 contains solution for https://adventofcode.com/2021/day/1 puzzle.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2021.String()
}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var (
		increasednum int
		prev         int
		idx          int
	)

	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("atoi: %w", err)
		}

		if idx == 0 {
			prev = n
			idx++

			continue
		}

		if n > prev {
			increasednum++
		}

		prev = n
		idx++
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	return strconv.Itoa(increasednum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
