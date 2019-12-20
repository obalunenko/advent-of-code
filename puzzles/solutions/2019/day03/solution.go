package day03

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

func init() {
	puzzleName, err := puzzles.MakeName("2019", "day03")
	if err != nil {
		panic(err)
	}

	puzzles.Register(puzzleName, solution{
		name: puzzleName,
	})
}

type solution struct {
	name string
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	res := make([]map[pos]bool, 0, 2)

	for scanner.Scan() {
		line := scanner.Text()

		w := makeWire()
		if err := w.run(line); err != nil {
			return "", errors.Wrap(err, "failed to run wire")
		}

		res = append(res, w.m)
	}

	cross := findCross(res[0], res[1])

	mds := make([]int, 0, len(cross))

	for _, p := range cross {
		md := p.manhattan()
		mds = append(mds, md)
	}

	sort.Ints(mds)

	return strconv.Itoa(mds[0]), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}

type wire struct {
	pos pos
	m   map[pos]bool
}

type pos struct {
	x int
	y int
}

func makeWire() wire {
	return wire{
		pos: pos{
			x: 0,
			y: 0,
		},
		m: make(map[pos]bool),
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
		w.m[w.pos] = true
	}
}

func (w *wire) down(n int) {
	for i := 0; i < n; i++ {
		w.pos.y--
		w.m[w.pos] = true
	}
}

func (w *wire) left(n int) {
	for i := 0; i < n; i++ {
		w.pos.x--
		w.m[w.pos] = true
	}
}

func (w *wire) right(n int) {
	for i := 0; i < n; i++ {
		w.pos.x++
		w.m[w.pos] = true
	}
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

func findCross(wm1 map[pos]bool, wm2 map[pos]bool) []pos {
	res := make([]pos, 0, len(wm1))

	for p := range wm1 {
		p := p

		if wm2[p] {
			res = append(res, p)
		}
	}

	return res
}
