// Package day06 contains solution for https://adventofcode.com/2021/day/6 puzzle.
package day06

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"

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
	return part1(input)
}

func (solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func part1(input io.Reader) (string, error) {
	states, err := parseSchoolFishesStates(input)
	if err != nil {
		return "", fmt.Errorf("parse school fishes states: %w", err)
	}

	sch := newSchool(80)
	sch.addElderFishes(states)

	<-sch.populate()

	fishes := len(sch.fishes)

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

func (s *school) addElderFishes(states []int) {
	s.globalWg.Add(len(states))

	for _, st := range states {
		f := newFish(s.globalWg, s, s.days, st)

		s.addFish(f)
	}
}

type school struct {
	globalWg *sync.WaitGroup
	mu       *sync.Mutex
	days     int
	fishes   []*fish
}

func (s *school) getFishes() []int {
	res := make([]int, 0, len(s.fishes))

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, f := range s.fishes {
		res = append(res, f.state)
	}

	return res
}

func (s *school) populate() chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.fishes {
		go s.fishes[i].reproduce()
	}

	ch := make(chan struct{})

	go func() {
		s.globalWg.Wait()

		close(ch)
	}()

	return ch
}

func newSchool(daysToReproduce int) *school {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	return &school{
		globalWg: &wg,
		mu:       &mu,
		days:     daysToReproduce,
		fishes:   make([]*fish, 0, 0),
	}
}

func (s *school) addFish(f *fish) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.fishes = append(s.fishes, f)
}

type fish struct {
	fishWg   *sync.WaitGroup
	daysLeft int
	state    int
	school   *school
}

func newFish(wg *sync.WaitGroup, s *school, daysLeft, currentState int) *fish {
	f := &fish{
		fishWg:   wg,
		daysLeft: daysLeft,
		state:    currentState,
		school:   s,
	}

	return f
}

func (f *fish) reproduce() {
	defer f.fishWg.Done()

	childWg := sync.WaitGroup{}

	var (
		idx       int
		prevState int
	)

	for f.daysLeft > 0 {
		if idx != 0 {
			f.state--
			f.daysLeft--
		}

		if prevState == 0 && idx != 0 {
			f.state = 6

			childWg.Add(1)

			child := newFish(&childWg, f.school, f.daysLeft, 8)

			f.school.addFish(child)

			go child.reproduce()
		}

		prevState = f.state
		idx++
	}

	childWg.Wait()
}
