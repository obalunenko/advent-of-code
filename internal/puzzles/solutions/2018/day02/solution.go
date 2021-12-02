// Package day02 contains solution for https://adventofcode.com/2018/day/2 puzzle.
package day02

import (
	"bufio"
	"errors"
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
	boxes, err := makeBoxesList(input)
	if err != nil {
		return "", fmt.Errorf("make boxes: %w", err)
	}

	const (
		two   = 2
		three = 3
	)

	var (
		twoCount, threeCount int
	)

	for i := range boxes {
		box := boxes[i]

		if hasNSameLetters(box, two) {
			twoCount++
		}

		if hasNSameLetters(box, three) {
			threeCount++
		}
	}

	checksum := twoCount * threeCount

	return strconv.Itoa(checksum), nil
}

var errNotFound = errors.New("not found")

func (s solution) Part2(input io.Reader) (string, error) {
	boxes, err := makeBoxesList(input)
	if err != nil {
		return "", fmt.Errorf("make boxes: %w", err)
	}

	// finds same boxes
	// find common letters
	boxesnum := len(boxes)

	var (
		box1, box2, common string
	)

loop:
	for i := 0; i <= boxesnum; i++ {
		for j := i + 1; j <= boxesnum-1; j++ {
			if hasNDiffLetters(boxes[i], boxes[j], 1) {
				box1, box2 = boxes[i], boxes[j]

				common = getCommonBoxesPart(box1, box2)

				break loop
			}
		}
	}

	if common == "" {
		return "", errNotFound
	}

	return common, nil
}

func getCommonBoxesPart(box1, box2 string) string {
	if len(box1) != len(box2) {
		return ""
	}

	br1, br2 := []rune(box1), []rune(box2)

	common := make([]rune, 0, len(br1))

	for i := 0; i < len(br1); i++ {
		a, b := br1[i], br2[i]
		if a == b {
			common = append(common, a)
		}
	}

	return string(common)
}

func hasNDiffLetters(box1, box2 string, n int) bool {
	if n < 0 {
		return false
	}

	if len(box1) != len(box2) {
		return false
	}

	var diff int

	for i := 0; i < len(box1); i++ {
		if box1[i] == box2[i] {
			continue
		}

		diff++
	}

	return diff == n
}

func makeBoxesList(input io.Reader) ([]string, error) {
	var boxes []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		boxes = append(boxes, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return boxes, nil
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
