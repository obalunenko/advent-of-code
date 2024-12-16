// Package day01 contains solution for https://adventofcode.com/2024/day/1 puzzle.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2024.String()
}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	l, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	slices.Sort(l.itemsA)
	slices.Sort(l.itemsB)

	var sum int

	for i := 0; i < len(l.itemsA); i++ {
		d := l.itemsA[i] - l.itemsB[i]
		if d < 0 {
			d = -d
		}

		sum += d
	}

	return strconv.Itoa(sum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	l, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	seenA := make(map[int]int)

	for _, a := range l.itemsA {
		seenA[a] = 0

		for _, b := range l.itemsB {
			if a == b {
				seenA[a]++
			}
		}
	}

	var sum int

	for _, a := range l.itemsA {
		sum += a * seenA[a]
	}

	return strconv.Itoa(sum), nil
}

type lists struct {
	itemsA []int
	itemsB []int
}

func parseInput(input io.Reader) (lists, error) {
	const (
		listsNum = 2
		listAIdx = 0
		listBIdx = 1
	)

	l := lists{
		itemsA: make([]int, 0),
		itemsB: make([]int, 0),
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "   ")
		if len(parts) != listsNum {
			return lists{}, fmt.Errorf("invalid input line: %s", line)
		}

		// Parse parts[0] and parts[1] to integers and append them to l.itemsA and l.itemsB respectively.
		a, err := strconv.Atoi(parts[listAIdx])
		if err != nil {
			return lists{}, fmt.Errorf("failed to parse int: %w", err)
		}

		b, err := strconv.Atoi(parts[listBIdx])
		if err != nil {
			return lists{}, fmt.Errorf("failed to parse int: %w", err)
		}

		l.itemsA = append(l.itemsA, a)

		l.itemsB = append(l.itemsB, b)
	}

	if scanner.Err() != nil {
		return lists{}, fmt.Errorf("scanner error: %w", scanner.Err())
	}

	return l, nil
}
