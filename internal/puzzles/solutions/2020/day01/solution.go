package day01

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day01"
	year       = "2020"
)

type solution struct {
	year string
	name string
}

func (s solution) Year() string {
	return s.year
}

func (s solution) Name() string {
	return s.name
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var expensereport = make(map[int]bool)

	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "", fmt.Errorf("[%s:%s]: part1: faied to parse int: %w", s.year, s.name, err)
		}

		expensereport[entry] = true
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("[%s:%s]: part1: scanner error: %w", s.year, s.name, err)
	}

	const (
		dest = 2020
	)

	var (
		a, b int
	)

	for e := range expensereport {
		a = e
		b = dest - e
		if expensereport[b] {
			break
		}
	}

	res := a * b

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var expensereport []int

	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "", fmt.Errorf("[%s:%s]: part2: faied to parse int: %w", s.year, s.name, err)
		}

		expensereport = append(expensereport, entry)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("[%s:%s]: part2: scanner error: %w", s.year, s.name, err)
	}

	sort.Ints(expensereport)

	const (
		dest = 2020
	)

	var (
		a, b, c int
		found   bool
	)

loop:
	for i := 0; i < len(expensereport)-2; i++ {
		a = expensereport[i]
		for n := i + 1; n < len(expensereport)-1; n++ {
			b = expensereport[n]
			for z := i + 2; z < len(expensereport); z++ {
				c = expensereport[z]

				sum := a + b + c
				if sum == dest && a != b && b != c {
					found = true

					break loop
				}
			}
		}

	}

	if !found {
		return "", fmt.Errorf("[%s:%s]: part2: answer not found", s.year, s.name)
	}

	res := a * b * c

	return strconv.Itoa(res), nil
}
