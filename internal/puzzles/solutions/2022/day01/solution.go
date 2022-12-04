// Package day01 contains solution for https://adventofcode.com/2022/day/1 puzzle.
package day01

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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
	list, err := makeElvesList(input)
	if err != nil {
		return "", err
	}

	res := list.backupSnackCalc()

	return strconv.Itoa(res), nil
}

func makeElvesList(input io.Reader) (elves, error) {
	scanner := bufio.NewScanner(input)

	var (
		list elves
		e    elve
		prev string
		line string
		n    int
		err  error
	)

	for scanner.Scan() {
		prev = line

		line = scanner.Text()
		if line == "" {
			list = append(list, e)

			e = elve{}

			continue
		}

		n, err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}

		e.food = append(e.food, n)
	}

	if prev == "" {
		list = append(list, e)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return list, nil
}

type elve struct {
	food  []int
	total int
}

func (e elve) String() string {
	return strconv.Itoa(e.totalCalories())
}

func (e elve) totalCalories() int {
	if e.total != 0 {
		return e.total
	}

	var sum int

	for i := range e.food {
		f := e.food[i]

		sum += f
	}

	e.total = sum

	return sum
}

type elves []elve

func (e elves) String() string {
	var resp string

	for _, e2 := range e {
		resp += fmt.Sprintln(e2.String())
	}

	return resp
}

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

func (e elves) backupSnackCalc() int {
	sort.Slice(e, func(i, j int) bool {
		return e[i].totalCalories() > e[j].totalCalories()
	})

	var sum int

	for i := 0; i < 3; i++ {
		sum += e[i].totalCalories()
	}

	return sum
}
