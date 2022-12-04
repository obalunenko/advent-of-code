// Package day01 contains solution for https://adventofcode.com/2022/day/1 puzzle.
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
	return puzzles.Year2022.String()
}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	list, err := makeElvesList(input)
	if err != nil {
		return "", err
	}

	res := list.maxTotalCalories()

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func makeElvesList(input io.Reader) (elves, error) {
	scanner := bufio.NewScanner(input)

	var list elves

	var e elve

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			list = append(list, e)

			e = elve{}

			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}

		e.food = append(e.food, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return list, nil
}

type elve struct {
	food []int
}

func (e elve) totalCalories() int {
	var sum int

	for i := range e.food {
		f := e.food[i]

		sum += f
	}

	return sum
}

type elves []elve

func (e elves) maxTotalCalories() int {
	var max int

	for _, el := range e {
		if max == 0 {
			max = el.totalCalories()

			continue
		}

		c := el.totalCalories()

		if c > max {
			max = c
		}
	}

	return max
}
