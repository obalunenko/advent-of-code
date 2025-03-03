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

	s.calcDistances(part1Cost)

	cost := s.minDistanceCost()

	return strconv.Itoa(cost), nil
}

func (solution) Part2(input io.Reader) (string, error) {
	crabs, err := getCrabs(input)
	if err != nil {
		return "", fmt.Errorf("get crabs: %w", err)
	}

	s := makeSwarm(crabs)

	s.calcDistances(part2Cost)

	cost := s.minDistanceCost()

	return strconv.Itoa(cost), nil
}

func getCrabs(input io.Reader) ([]int, error) {
	crabs, err := utils.ParseInts(input, ",")
	if err != nil {
		return nil, fmt.Errorf("parse int slice from reader: %w", err)
	}

	return crabs, nil
}

const (
	undef = -99999
)

func makeMatrix(crabs []int) [][]int {
	const (
		header = 1
	)

	sort.Ints(crabs)

	cnum := len(crabs)

	maxC := crabs[cnum-1]

	matrix := make([][]int, cnum+header)

	// matrix[i][j].
	// 	i - crabs; j - all positions from 0 to maxC
	for i := 0; i < cnum+header; i++ {
		matrix[i] = make([]int, maxC+header+1)

		if i == 0 {
			matrix[i][0] = undef

			for j := 1; j <= maxC+header; j++ {
				matrix[i][j] = j - 1
			}

			continue
		}

		matrix[i][0] = crabs[i-1]
	}

	return matrix
}

type swarm struct {
	crabsNum     int
	distancesNum int
	crabsMatrix  [][]int
}

func (s swarm) getMatrixILen() int {
	return s.crabsNum + 1
}

func (s swarm) getMatrixJLen() int {
	return s.distancesNum + 1
}

func makeSwarm(crabs []int) swarm {
	matrix := makeMatrix(crabs)

	crabsNum := len(matrix)

	distNum := len(matrix[0])

	return swarm{
		crabsNum:     crabsNum - 1,
		distancesNum: distNum - 1,
		crabsMatrix:  matrix,
	}
}

type fuelCostFunc func(p int) int

func part1Cost(p int) int {
	return p
}

//nolint:mnd // Formula is not a magic number.
func part2Cost(p int) int {
	// formula a_{n}=a_{1}+(n-1)d
	an := 1 + 1*(p-1)

	// formula s_{n}=(a_{1}+a_{n})/2*n
	s := ((1 + an) * p) / 2

	return s
}

func (s *swarm) calcDistances(cost fuelCostFunc) {
	for i := 1; i < s.getMatrixILen(); i++ {
		for j := 1; j < s.getMatrixJLen(); j++ {
			p := s.crabsMatrix[i][0] - s.crabsMatrix[0][j]
			if p < 0 {
				p *= -1
			}

			s.crabsMatrix[i][j] = cost(p)
		}
	}
}

func (s swarm) minDistanceCost() int {
	return minDistanceCost(s.crabsMatrix)
}

func minDistanceCost(matrix [][]int) int {
	var minC int

	ilen := len(matrix)
	jlen := len(matrix[0])

	for j := 1; j < jlen; j++ {
		var f int

		for i := 1; i < ilen; i++ {
			f += matrix[i][j]
		}

		if j == 1 {
			minC = f
		}

		if f < minC {
			minC = f
		}
	}

	return minC
}
