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

func (s solution) Name() string {
	return puzzles.Day01.String()
}

func (s solution) Year() string {
	return puzzles.Year2015.String()
}

func init() {
	puzzles.Register(solution{})
}

func (s solution) Part1(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)

	e := newElevator()

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", fmt.Errorf("read rune: %w", err)
		}

		e.Move(move(r))
	}

	return strconv.Itoa(int(e.Floor())), nil
}

func (s solution) Part2(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)

	e := newElevator()

	var pos int

	for e.Floor() != basement {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", fmt.Errorf("read rune: %w", err)
		}

		pos++

		e.Move(move(r))
	}

	return strconv.Itoa(pos), nil
}
