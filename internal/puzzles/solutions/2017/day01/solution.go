// Package day01 contains solution for https://adventofcode.com/2017/day/1 puzzle.
package day01

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Year() string {
	return puzzles.Year2017.String()
}

func init() {
	puzzles.Register(solution{})
}

func (s solution) Part1(in io.Reader) (string, error) {
	return part1(in)
}

func (s solution) Part2(in io.Reader) (string, error) {
	return part2(in)
}

func part1(in io.Reader) (string, error) {
	list, err := makeList(in)
	if err != nil {
		return "", fmt.Errorf("make list: %w", err)
	}

	shift := 1

	itr := newIterator(list, shift, true)

	sum := itr.Sum()

	return strconv.Itoa(sum), nil
}

func part2(in io.Reader) (string, error) {
	list, err := makeList(in)
	if err != nil {
		return "", fmt.Errorf("make list: %w", err)
	}

	shift := len(list) / 2

	itr := newIterator(list, shift, true)

	sum := itr.Sum()

	return strconv.Itoa(sum), nil
}

func makeList(in io.Reader) ([]int, error) {
	const newline = '\n'

	var list []int

	reader := bufio.NewReader(in)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("read rune: %w", err)
		}

		if r == newline {
			continue
		}

		n, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, fmt.Errorf("strconv atoi: %w", err)
		}

		list = append(list, n)
	}

	return list, nil
}

type iterator struct {
	list       []int
	shift      int
	isCircular bool
}

func newIterator(list []int, shift int, isCircular bool) iterator {
	return iterator{
		list:       list,
		shift:      shift,
		isCircular: isCircular,
	}
}

func (i iterator) Sum() int {
	var (
		cursorStart, cursorEnd int
		sum                    int
	)

	listlen := len(i.list)
	rightBound := listlen - 1
	lastidx := rightBound

	if !i.isCircular {
		lastidx -= i.shift
	}

	for cursorStart <= lastidx {
		cursorEnd = cursorStart

		cursorEnd += i.shift
		if i.isCircular {
			if cursorEnd > rightBound {
				cursorEnd -= listlen
			}
		}

		x, y := i.list[cursorStart], i.list[cursorEnd]
		if x == y {
			sum += x
		}

		cursorStart++
	}

	return sum
}
