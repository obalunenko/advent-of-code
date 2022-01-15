// Package day03 contains solution for https://adventofcode.com/2015/day/3 puzzle.
package day03

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

func (solution) Day() string {
	return puzzles.Day03.String()
}

func (solution) Year() string {
	return puzzles.Year2015.String()
}

func (solution) Part1(input io.Reader) (string, error) {
	reader := bufio.NewReader(input)

	s := newSanta()

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", fmt.Errorf("read rune: %w", err)
		}

		m := string(r)

		if r == '\n' {

			continue
		}

		if err = s.visit(m); err != nil {
			return "", fmt.Errorf("visit: %w", err)
		}
	}

	visited := s.housesVisited()

	return strconv.Itoa(visited), nil
}

func (solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

type grid struct {
	x, y int
}

func newGrid() grid {
	return grid{
		x: 0,
		y: 0,
	}
}

var errUnknownDirecton = errors.New("unknown direction")

func (g *grid) move(m string) error {
	const (
		north = "^"
		south = "v"
		east  = ">"
		west  = "<"
	)

	switch m {

	case north:
		g.y++
	case south:
		g.y--
	case west:
		g.x--
	case east:
		g.x++
	default:
		return errUnknownDirecton
	}

	return nil
}

type santa struct {
	location grid
	visited  map[grid]int
}

func newSanta() santa {
	startLoc := newGrid()

	visited := make(map[grid]int)
	visited[startLoc]++

	return santa{
		location: startLoc,
		visited:  visited,
	}
}

func (s *santa) visit(address string) error {
	if err := s.location.move(address); err != nil {
		return fmt.Errorf("failed to visit address: %w", err)
	}

	s.visited[s.location]++

	return nil
}

func (s santa) housesVisited() int {
	return len(s.visited)
}
