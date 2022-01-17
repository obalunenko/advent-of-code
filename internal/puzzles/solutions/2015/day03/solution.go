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
	addresses, err := makeAddressesList(input)
	if err != nil {
		return "", fmt.Errorf("make addresses: %w", err)
	}

	delivery := newSantaDelivery([]deliveryman{newSanta()})

	if err = delivery.deliver(addresses); err != nil {
		return "", fmt.Errorf("deliver: %w", err)
	}

	visited := delivery.housesVisited()

	return strconv.Itoa(visited), nil
}

func (solution) Part2(input io.Reader) (string, error) {
	addresses, err := makeAddressesList(input)
	if err != nil {
		return "", fmt.Errorf("make addresses: %w", err)
	}

	delivery := newSantaDelivery([]deliveryman{newSanta(), newSanta()})

	if err = delivery.deliver(addresses); err != nil {
		return "", fmt.Errorf("deliver: %w", err)
	}

	visited := delivery.housesVisited()

	return strconv.Itoa(visited), nil
}

func makeAddressesList(input io.Reader) ([]string, error) {
	reader := bufio.NewReader(input)

	var list []string

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("read rune: %w", err)
		}

		if r == '\n' {
			continue
		}

		list = append(list, string(r))
	}

	return list, nil
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

var errUnknownDirection = errors.New("unknown direction")

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
		return errUnknownDirection
	}

	return nil
}

type santaDelivery struct {
	santas []deliveryman
}

func newSantaDelivery(deliverymen []deliveryman) santaDelivery {
	return santaDelivery{
		santas: deliverymen,
	}
}

func (s *santaDelivery) deliver(addresses []string) error {
	for i, address := range addresses {
		snt := s.santas[0]

		if i%2 == 0 && len(s.santas) == 2 {
			snt = s.santas[1]
		}

		if err := snt.visit(address); err != nil {
			return err
		}
	}

	return nil
}

func (s *santaDelivery) housesVisited() int {
	total := make(map[grid]int)

	for i := range s.santas {
		snt := s.santas[i]

		for address, count := range snt.housesVisited() {
			total[address] += count
		}
	}

	return len(total)
}

type deliveryman interface {
	visit(address string) error
	housesVisited() map[grid]int
}

type santa struct {
	location grid
	visited  map[grid]int
}

func newSanta() *santa {
	startLoc := newGrid()

	visited := make(map[grid]int)
	visited[startLoc]++

	return &santa{
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

func (s santa) housesVisited() map[grid]int {
	return s.visited
}
