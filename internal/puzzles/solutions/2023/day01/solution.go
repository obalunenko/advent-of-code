// Package day01 contains solution for https://adventofcode.com/2023/day/1 puzzle.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
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
	sum, err := calibrate(input, nil)
	if err != nil {
		return "", fmt.Errorf("calibrating: %w", err)
	}

	return strconv.Itoa(sum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	sum, err := calibrate(input, digitsDict)
	if err != nil {
		return "", fmt.Errorf("calibrating: %w", err)
	}

	return strconv.Itoa(sum), nil
}

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var digitsDict = map[string]int{
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}

func calibrate(input io.Reader, dictionary map[string]int) (int, error) {
	scanner := bufio.NewScanner(input)

	var values []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		value, err := extractNumberFromLine(line, dictionary)
		if err != nil {
			return 0, fmt.Errorf("extracting number from line %q: %w", line, err)
		}

		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("reading input: %w", err)
	}

	var sum int

	for _, v := range values {
		sum += v
	}

	return sum, nil
}

func extractNumberFromLine(line string, dict map[string]int) (int, error) {
	first, last := -1, -1

	var word string

	for _, c := range line {
		if !unicode.IsDigit(c) {
			word += string(c)

			if d, ok := getDigitFromWord(word, dict); ok {
				if first == -1 {
					first = d
				} else {
					last = d
				}

				word = word[len(word)-1:]
			}

			continue
		}

		word = ""

		if first == -1 {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, fmt.Errorf("failed to convert %q to int: %w", string(c), err)
			}

			first = d

			continue
		}

		d, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, fmt.Errorf("failed to convert %q to int: %w", string(c), err)
		}

		last = d
	}

	if last == -1 {
		last = first
	}

	value, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
	if err != nil {
		return 0, fmt.Errorf("failed to convert %d%d to int: %w", first, last, err)
	}

	return value, nil
}

func getDigitFromWord(word string, dict map[string]int) (int, bool) {
	if word == "" {
		return -1, false
	}

	if len(dict) == 0 {
		return -1, false
	}

	for s, i := range digitsDict {
		if strings.Contains(word, s) {
			return i, true
		}
	}

	return -1, false
}
