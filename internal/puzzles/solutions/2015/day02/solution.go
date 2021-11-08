// Package day02 contains solution for https://adventofcode.com/2015/day/2 puzzle.
package day02

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Year() string {
	return puzzles.Year2015.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	return part1(input)
}

func (s solution) Part2(input io.Reader) (string, error) {
	return part2(input)
}

func part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var res int

	for scanner.Scan() {
		line := scanner.Text()

		b, err := boxFromDimensions(line)
		if err != nil {
			return "", fmt.Errorf("failed to make box: %w", err)
		}

		res += b.surfaceWithExtra()
	}

	return strconv.Itoa(res), nil
}

func part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var res int

	for scanner.Scan() {
		line := scanner.Text()

		b, err := boxFromDimensions(line)
		if err != nil {
			return "", fmt.Errorf("failed to make box: %w", err)
		}

		res += b.ribbon()
	}

	return strconv.Itoa(res), nil
}

type rectangle struct {
	width  int
	length int
}

func (r rectangle) square() int {
	return r.width * r.length
}

type cuboid struct {
	width  int
	length int
	height int
}

func (c cuboid) faces() []rectangle {
	return []rectangle{
		{
			width:  c.width,
			length: c.height,
		},
		{
			width:  c.height,
			length: c.length,
		},
		{
			width:  c.length,
			length: c.width,
		},
	}
}

type box struct {
	cuboid
}

// 2x4x7, where x is delimiter.
func boxFromDimensions(dimensions string) (box, error) {
	const (
		delim  = "x"
		dimNum = 3

		widthPos  = 0
		heightPos = 1
		lengthPos = 2
	)

	dims := strings.Split(dimensions, delim)
	if len(dims) != dimNum {
		return box{}, fmt.Errorf("invaid dimensions")
	}

	var sides = make([]int, 0, 3)

	for _, s := range dims {
		d, err := strconv.Atoi(s)
		if err != nil {
			return box{}, fmt.Errorf("atoi dimension: %w", err)
		}

		sides = append(sides, d)
	}

	w := sides[widthPos]
	h := sides[heightPos]
	l := sides[lengthPos]

	b := newBox(w, h, l)

	return b, nil
}

func newBox(width, height, length int) box {
	return box{
		cuboid: cuboid{
			width:  width,
			length: length,
			height: height,
		},
	}
}

func (b box) surfaceWithExtra() int {
	var (
		// the smallest surface area used an extra paper.
		extra    int
		surfarea int
	)

	for _, f := range b.faces() {
		s := f.square()

		if extra == 0 {
			extra = s
		}

		if s < extra {
			extra = s
		}

		surfarea += s
	}

	// each face meets twice.
	const mod = 2

	surfarea *= mod

	return surfarea + extra
}

func (b box) wrapRibbon() int {
	var sides = []int{
		b.height, b.width, b.length,
	}

	sort.Ints(sides)

	const (
		// each side meets twice
		mod    = 2
		first  = 0
		second = 1
	)

	s := sides[first]*mod + sides[second]*mod

	return s
}

func (b box) bowRibbon() int {
	return b.width * b.height * b.length
}

func (b box) ribbon() int {
	return b.bowRibbon() + b.wrapRibbon()
}
