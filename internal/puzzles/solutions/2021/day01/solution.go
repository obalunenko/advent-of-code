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
	list, err := makeMeasurementsList(input)
	if err != nil {
		return "", fmt.Errorf("make measurements list: %w", err)
	}

	const (
		shift      = 1
		windowSize = 1
	)

	increasednum := findIncreased(list, shift, windowSize)

	return strconv.Itoa(increasednum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	list, err := makeMeasurementsList(input)
	if err != nil {
		return "", fmt.Errorf("make measurements list: %w", err)
	}

	const (
		shift      = 1
		windowSize = 3
	)

	increasednum := findIncreased(list, shift, windowSize)

	return strconv.Itoa(increasednum), nil
}

func findIncreased(list []int, shift, window int) int {
	var increadsed int

	for i := 0; i <= len(list)-window; i += shift {
		if i == 0 {
			continue
		}

		var m1, m2 int

		k := i
		for j := window; j > 0; j-- {
			m2 += list[k]
			m1 += list[k-shift]

			k++
		}

		if m2 > m1 {

			increadsed++
		}
	}

	return increadsed
}

func makeMeasurementsList(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)

	var measurements []int

	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}

		measurements = append(measurements, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return measurements, nil
}
