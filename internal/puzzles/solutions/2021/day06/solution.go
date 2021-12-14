// Package day06 contains solution for https://adventofcode.com/2021/day/6 puzzle.
package day06

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (solution) Day() string {
	return puzzles.Day06.String()
}

func (solution) Year() string {
	return puzzles.Year2021.String()
}

func (solution) Part1(input io.Reader) (string, error) {
	return observeFishSchool(input, 80)
}

func (solution) Part2(input io.Reader) (string, error) {
	return observeFishSchool(input, 256)
}

func observeFishSchool(input io.Reader, days int) (string, error) {
	states, err := parseSchoolFishesStates(input)
	if err != nil {
		return "", fmt.Errorf("parse school fishes states: %w", err)
	}

	sch := newSchool(days)
	sch.addElderFishes(states)

	sch.populate()

	fishes := sch.getFishes()

	return strconv.Itoa(fishes), nil
}

func parseSchoolFishesStates(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)

	var res []int

	for scanner.Scan() {
		line := scanner.Text()

		states := strings.Split(line, ",")
		for _, s := range states {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("parse fish state: %w", err)
			}

			res = append(res, n)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return res, nil
}

func (s *school) addElderFishes(fishes []int) {
	for _, st := range fishes {
		s.fishes[st]++
	}
}

type school struct {
	days   int
	fishes map[int]int
}

func (s *school) getFishes() int {
	var res int

	for _, f := range s.fishes {
		res += f
	}

	return res
}

func (s *school) populate() {
	for d := s.days; d > 0; d-- {
		for i := 0; i <= 8; i++ {
			s.fishes[i-1] += s.fishes[i]
			s.fishes[i] = 0
		}

		s.fishes[8] += s.fishes[-1]
		s.fishes[6] += s.fishes[-1]
		s.fishes[-1] = 0
	}
}

func newSchool(daysToReproduce int) *school {

	return &school{
		days:   daysToReproduce,
		fishes: make(map[int]int),
	}
}
