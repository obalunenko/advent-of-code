// Package day01 contains solution for https://adventofcode.com/2023/day/1 puzzle.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"unicode"

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

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var values []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		first, last := -1, -1

		for _, c := range line {
			if !unicode.IsDigit(c) {
				continue
			}

			if first == -1 {
				d, err := strconv.Atoi(string(c))
				if err != nil {
					return "", fmt.Errorf("failed to convert %s to int: %w", string(c), err)
				}

				first = d
			}

			if first != -1 {
				d, err := strconv.Atoi(string(c))
				if err != nil {
					return "", fmt.Errorf("failed to convert %s to int: %w", string(c), err)
				}

				last = d
			}
		}

		if last == -1 {
			last = first
		}

		value, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			return "", fmt.Errorf("failed to convert %d%d to int: %w", first, last, err)
		}

		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("reading input: %w", err)
	}

	var sum int

	for _, v := range values {
		sum += v
	}

	return strconv.Itoa(sum), nil
}

func (s solution) Part2(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
