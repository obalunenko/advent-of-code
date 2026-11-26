// Package day02 contains solution for https://adventofcode.com/2024/day/2 puzzle.
package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/common/utils"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2024.String()
}

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	isSafe := func(line []int) bool {
		var asc, desc bool

		for i, val := range line {
			if i == 0 {
				continue
			}

			if line[i-1] < val {
				if desc {
					return false
				}

				asc = true
			}

			if line[i-1] > val {
				if asc {
					return false
				}

				desc = true
			}

			diff := val - line[i-1]
			if diff < 0 {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				return false
			}
		}

		return true
	}

	var safeCount int

	for scanner.Scan() {
		line := scanner.Bytes()

		numbers, err := utils.ParseInts(bytes.NewReader(line), " ")
		if err != nil {
			return "", fmt.Errorf("failed to parse input line: %w", err)
		}

		if isSafe(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strconv.Itoa(safeCount), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	isSafe := func(line []int) bool {
		var asc, desc bool

		removedCount := 0

		mayRemove := func(i int) bool {
			if removedCount != 0 {
				return false
			}

			removedCount++

			return true
		}

		for i := range line {
			if i == len(line)-1 {

			}

			if line[i-1] < val {
				if desc {
					return false
				}

				asc = true
			}

			if line[i-1] > val {
				if asc {
					return false
				}

				desc = true
			}

			diff := val - line[i-1]
			if diff < 0 {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				if !mayRemove(i) {
					diff = line[i-1] - line[i+1]
					if diff < 0 {
						diff = -diff
					}

					if diff < 1 || diff > 3 {
						return false
					}

					continue
				}

				return false
			}
		}

		return true
	}

	var safeCount int

	for scanner.Scan() {
		line := scanner.Bytes()

		numbers, err := utils.ParseInts(bytes.NewReader(line), " ")
		if err != nil {
			return "", fmt.Errorf("failed to parse input line: %w", err)
		}

		if isSafe(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strconv.Itoa(safeCount), nil
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)

	ret = append(ret, s[:index]...)

	return append(ret, s[index+1:]...)
}
