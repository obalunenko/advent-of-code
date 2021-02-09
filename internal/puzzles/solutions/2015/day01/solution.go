package day01

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day01"
	year       = "2015"
)

type solution struct {
	year string
	name string
}

const (
	up   = "("
	down = ")"
)

func (s solution) Part1(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)

	var floor int

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", fmt.Errorf("read rune: %w", err)
		}

		s := string(r)

		switch s {
		case up:
			floor++
		case down:
			floor--
		}
	}

	return strconv.Itoa(floor), nil
}

func (s solution) Part2(_ io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}

func (s solution) Year() string {
	return s.year
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}
