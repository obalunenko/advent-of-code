// Package day03 solves https://adventofcode.com/2019/day/3
package day03

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day03"
	year       = "2019"
)

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

type solution struct {
	year string
	name string
}

func (s solution) Year() string {
	return s.year
}

func (s solution) Part1(input io.Reader) (string, error) {
	wires, err := runWires(input)
	if err != nil {
		return "", err
	}

	cross := findCross(wires[0], wires[1])

	mds := make([]int, 0, len(cross))

	for _, p := range cross {
		md := p.manhattan()
		mds = append(mds, md)
	}

	sort.Ints(mds)

	return strconv.Itoa(mds[0]), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	wires, err := runWires(input)
	if err != nil {
		return "", err
	}

	cross := findCross(wires[0], wires[1])

	stps := make([]int, 0, len(cross))

	for _, c := range cross {
		sum := wires[0][c] + wires[1][c]
		stps = append(stps, sum)
	}

	sort.Ints(stps)

	return strconv.Itoa(stps[0]), nil
}

func (s solution) Name() string {
	return s.name
}

type wire struct {
	pos  pos
	step stepper
	m    map[pos]int
}

type pos struct {
	x int
	y int
}

type stepper struct {
	steps int
}

func (s *stepper) add() {
	s.steps++
}

func (s stepper) get() int {
	return s.steps
}

func makeWire() wire {
	return wire{
		pos: pos{
			x: 0,
			y: 0,
		},
		m: make(map[pos]int),
	}
}

const (
	moveUp    = "U"
	moveDown  = "D"
	moveRight = "R"
	moveLeft  = "L"
)

func (w *wire) run(input string) error {
	moves := strings.Split(input, ",")
	for _, m := range moves {
		move := m

		act := string(move[0])

		steps, err := strconv.Atoi(move[1:])
		if err != nil {
			return errors.New("failed to parse steps")
		}

		switch act {
		case moveDown:
			w.down(steps)
		case moveUp:
			w.up(steps)
		case moveLeft:
			w.left(steps)
		case moveRight:
			w.right(steps)
		}
	}

	return nil
}

func (w *wire) up(n int) {
	for i := 0; i < n; i++ {
		w.pos.y++
		w.storePosition()
	}
}

func (w *wire) down(n int) {
	for i := 0; i < n; i++ {
		w.pos.y--
		w.storePosition()
	}
}

func (w *wire) left(n int) {
	for i := 0; i < n; i++ {
		w.pos.x--
		w.storePosition()
	}
}

func (w *wire) right(n int) {
	for i := 0; i < n; i++ {
		w.pos.x++
		w.storePosition()
	}
}

func (w *wire) storePosition() {
	w.step.add()
	w.m[w.pos] = w.step.get()
}

func (p pos) manhattan() int {
	x, y := p.x, p.y

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

func findCross(wm1 map[pos]int, wm2 map[pos]int) []pos {
	res := make([]pos, 0, len(wm1))

	for p := range wm1 {
		p := p

		if _, exist := wm2[p]; exist {
			res = append(res, p)
		}
	}

	return res
}

func runWires(input io.Reader) ([]map[pos]int, error) {
	scanner := bufio.NewScanner(input)
	res := make([]map[pos]int, 0, 2)

	for scanner.Scan() {
		line := scanner.Text()

		w := makeWire()
		if err := w.run(line); err != nil {
			return nil, errors.Wrap(err, "failed to run wire")
		}

		res = append(res, w.m)
	}

	return res, nil
}
