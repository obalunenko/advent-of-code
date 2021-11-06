// Package day04 contains solution for https://adventofcode.com/2019/day/4 puzzle.
package day04

import (
	"bytes"
	"errors"
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

func (s solution) Year() string {
	return puzzles.Year2019.String()
}

func (s solution) Day() string {
	return puzzles.Day04.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	return run(input, isPasswordPart1)
}

func (s solution) Part2(input io.Reader) (string, error) {
	return run(input, isPasswordPart2)
}

func run(input io.Reader, criteria isPwdFunc) (string, error) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(input); err != nil {
		return "", fmt.Errorf("failed to read: %w", err)
	}

	const limitsNum = 2

	raw := strings.TrimSpace(buf.String())

	limits := strings.Split(raw, "-") // should be 2: low and high
	if len(limits) != limitsNum {
		return "", errors.New("invalid number of limits")
	}

	passwords, err := findPasswords(limits[0], limits[1], criteria)
	if err != nil {
		return "", fmt.Errorf("failed to find passwords: %w", err)
	}

	return strconv.Itoa(passwords), nil
}

type isPwdFunc func(n int) bool

func findPasswords(low, high string, criteria isPwdFunc) (int, error) {
	lowd, err := strconv.Atoi(low)
	if err != nil {
		return -1, fmt.Errorf("failed to convert low to int: %w", err)
	}

	highd, err := strconv.Atoi(high)
	if err != nil {
		return -1, fmt.Errorf("failed to convert high to int: %w", err)
	}

	pwds := make([]int, 0, highd-lowd)

	for i := lowd; i <= highd; i++ {
		if criteria(i) {
			pwds = append(pwds, i)
		}
	}

	return len(pwds), nil
}

func isIncreasing(n int) bool {
	nmbs := intToSlice(n)

	for i := 1; i <= len(nmbs)-1; i++ {
		if nmbs[i] < nmbs[i-1] {
			return false
		}
	}

	return true
}

func hasRepeated(n int) bool {
	nmbs := intToSlice(n)

	var hasRepeated bool

	for i := 1; i <= len(nmbs)-1; i++ {
		if nmbs[i] == nmbs[i-1] {
			hasRepeated = true
		}
	}

	return hasRepeated
}

func hasRepeatedWithDouble(n int) bool {
	nmbs := intToSlice(n)

	repeated := make(map[int]int)

	for i := 1; i <= len(nmbs)-1; i++ {
		if nmbs[i] == nmbs[i-1] {
			repeated[nmbs[i]]++
		}
	}

	if len(repeated) == 0 {
		return false
	}

	var hasDouble bool

	for i := 1; i < 10; i++ {
		if n, ok := repeated[i]; ok {
			if n == 1 {
				hasDouble = true
			}
		}
	}

	return hasDouble
}

func isPasswordPart1(n int) bool {
	return isIncreasing(n) && hasRepeated(n)
}

func isPasswordPart2(n int) bool {
	return isIncreasing(n) && hasRepeatedWithDouble(n)
}

func intToSlice(n int) [6]int {
	return [6]int{
		(n % 1000000) / 100000,
		(n % 100000) / 10000,
		(n % 10000) / 1000,
		(n % 1000) / 100,
		(n % 100) / 10,
		n % 10,
	}
}
