// Package day02 contains solution for https://adventofcode.com/2017/day/2 puzzle.
package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Year() string {
	return puzzles.Year2017.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	var f checksumFunc = func(row []string) (int, error) {
		var min, max int

		for i, number := range row {
			d, err := strconv.Atoi(number)
			if err != nil {
				return 0, fmt.Errorf("atoi: %w", err)
			}

			if i == 0 {
				min, max = d, d
			}

			if d < min {
				min = d
			}

			if d > max {
				max = d
			}
		}

		return max - min, nil
	}

	return findChecksum(input, f)
}

func (s solution) Part2(input io.Reader) (string, error) {
	var f checksumFunc = func(row []string) (int, error) {
		numbers := make([]int, 0, len(row))

		for _, n := range row {
			d, err := strconv.Atoi(n)
			if err != nil {
				return 0, fmt.Errorf("atoi: %w", err)
			}

			numbers = append(numbers, d)
		}

		var result int

	loop:
		for i := 0; i < len(numbers); i++ {
			d1 := numbers[i]
			for j := i + 1; j < len(numbers); j++ {
				d2 := numbers[j]

				var a, b = d1, d2

				if a < b {
					a, b = b, a
				}

				if a%b == 0 {
					result = a / b

					break loop
				}
			}
		}

		return result, nil
	}

	return findChecksum(input, f)
}

type checksumFunc func(row []string) (int, error)

func findChecksum(spreadsheet io.Reader, f checksumFunc) (string, error) {
	scanner := bufio.NewScanner(spreadsheet)

	var checksum int

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.ReplaceAll(line, "\t", " ")

		numbers := strings.Split(line, " ")

		n, err := f(numbers)
		if err != nil {
			return "", err
		}

		checksum += n
	}

	return strconv.Itoa(checksum), nil

}
