// Package day07 contains solution for https://adventofcode.com/2021/day/7 puzzle.
package day07

import (
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/common/utils"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (solution) Day() string {
	return puzzles.Day07.String()
}

func (solution) Year() string {
	return puzzles.Year2021.String()
}

func (solution) Part1(input io.Reader) (string, error) {
	crabs, err := getCrabs(input)
	if err != nil {
		return "", fmt.Errorf("get crabs: %w", err)
	}

	s := makeSwarm(crabs)

	s.calcDistances()

	cost := s.minDistanceCost()

	return strconv.Itoa(cost), nil
}

func (solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func getCrabs(input io.Reader) ([]int, error) {
	crabs, err := utils.ParseInts(input, ",")
	if err != nil {
		return nil, fmt.Errorf("parse int slice from reader: %w", err)
	}

	sort.Ints(crabs)

	return crabs, nil
}

const (
	undef = -99999
)

func makeMatrix(crabs []int) [][]int {
	const (
		header    = 1
		headerPos = 0 // matrix[y][x] - where 0 - x
	)

	cnum := len(crabs) + header

	matrix := make([][]int, cnum)

	// matrix[i][j]
	for i := 0; i < cnum; i++ {
		matrix[i] = make([]int, cnum)

		for j := 0; j < cnum; j++ {
			switch {
			case i == headerPos && j == headerPos:
				matrix[i][j] = undef
			case i == headerPos:
				matrix[i][j] = crabs[j-1]
			case j == headerPos:
				matrix[i][j] = crabs[i-1]
			}
		}
	}

	return matrix
}

type swarm struct {
	crabsMatrix [][]int
}

func makeSwarm(crabs []int) swarm {
	return swarm{
		crabsMatrix: makeMatrix(crabs),
	}
}

func (s *swarm) calcDistances() {
	for i := 1; i < len(s.crabsMatrix); i++ {
		for j := 1; j < len(s.crabsMatrix); j++ {
			p := s.crabsMatrix[i][0] - s.crabsMatrix[0][j]
			if p < 0 {
				p *= -1
			}

			s.crabsMatrix[i][j] = p
		}
	}
}

func (s swarm) minDistanceCost() int {
	var min int

	for i := 1; i < len(s.crabsMatrix); i++ {
		var f int

		for j := 1; j < len(s.crabsMatrix); j++ {
			f += s.crabsMatrix[i][j]
		}

		if i == 1 {
			min = f
		}

		if f < min {
			min = f
		}
	}

	return min
}
