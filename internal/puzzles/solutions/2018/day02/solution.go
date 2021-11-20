// Package day02 contains solution for https://adventofcode.com/2018/day/2 puzzle.
package day02

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

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Year() string {
	return puzzles.Year2018.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	const (
		two   = 2
		three = 3
	)

	var (
		twoCount, threeCount int
	)

	for scanner.Scan() {
		line := scanner.Text()

		if hasNSameLetters(line, two) {
			twoCount++
		}

		if hasNSameLetters(line, three) {
			threeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	checksum := twoCount * threeCount

	return strconv.Itoa(checksum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func hasNSameLetters(s string, n int) bool {
	if n <= 0 {
		return false
	}

	seen := make(map[rune]int)

	for _, c := range s {
		seen[c]++
	}

	for _, i := range seen {
		if i == n {
			return true
		}
	}

	return false
}
