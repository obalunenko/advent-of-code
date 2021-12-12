// Package day05 contains solution for https://adventofcode.com/2021/day/5 puzzle.
package day05

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2021.String()
}

func (s solution) Day() string {
	return puzzles.Day05.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	lines, err := getLines(input)
	if err != nil {
		return "", fmt.Errorf("get lines: %w", err)
	}

	lines = filterLines(lines, part1Filter)

	d := drawDiagram(lines)

	zones := d.dangerZones(part1DangerZone)

	return strconv.Itoa(zones), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

type position struct {
	x, y int
}

type line struct {
	start position
	end   position
}

type diagram struct {
	data [][]int
}

type dangerFunc func(n int) bool

func part1DangerZone(n int) bool {
	return n > 1
}

func (d diagram) dangerZones(f dangerFunc) int {
	var zones int

	for _, xs := range d.data {
		for _, x := range xs {
			if f(x) {
				zones++
			}
		}
	}

	return zones
}

func (d diagram) String() string {
	const (
		empty   = "."
		newline = "\n"
	)

	var res string

	last := len(d.data) - 1

	for i := 0; i <= last; i++ {
		xs := d.data[i]

		for _, x := range xs {
			if x == 0 {
				res += empty

				continue
			}

			res += strconv.Itoa(x)
		}

		if i != last {
			res += newline
		}
	}

	return res
}

func (d *diagram) draw(lines []line) {
	for _, l := range lines {
		if l.start.x == l.end.x {
			d.drawY(l)
		}

		if l.start.y == l.end.y {
			d.drawX(l)
		}
	}
}

func (d *diagram) drawX(l line) {
	y := l.start.y

	x1, x2 := l.start.x, l.end.x
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	for i := x1; i <= x2; i++ {
		d.data[y][i]++
	}
}

func (d *diagram) drawY(l line) {
	x := l.start.x

	y1, y2 := l.start.y, l.end.y
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for i := y1; i <= y2; i++ {
		d.data[i][x]++
	}
}

func newDiagram(maxX, maxY int) diagram {
	res := make([][]int, maxY+1)

	for i := 0; i < maxY+1; i++ {
		res[i] = make([]int, maxX+1)
	}

	return diagram{
		data: res,
	}
}

func drawDiagram(lines []line) diagram {

	// get max x,y
	bounds := getBounds(lines)

	// allocate
	d := newDiagram(bounds.x, bounds.y)

	// draw
	d.draw(lines)

	return d
}

func getBounds(lines []line) position {
	var (
		maxX, maxY int
	)

	for _, l := range lines {
		if l.start.x > maxX {
			maxX = l.start.x
		}

		if l.start.y > maxY {
			maxY = l.start.y
		}

		if l.end.x > maxX {
			maxX = l.end.x
		}

		if l.end.y > maxY {
			maxY = l.end.y
		}
	}

	return position{
		x: maxX,
		y: maxY,
	}
}

func getLines(input io.Reader) ([]line, error) {
	scanner := bufio.NewScanner(input)

	var lines []line

	const (
		startpos = 0
		endpos   = 1
	)

	for scanner.Scan() {
		l := scanner.Text()

		coordinates, err := parseLine(l)
		if err != nil {
			return nil, fmt.Errorf("get numbers: %w", err)
		}

		lines = append(lines, line{
			start: coordinates[startpos],
			end:   coordinates[endpos],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return lines, nil
}

var reg = regexp.MustCompile(`(?s)\d+,\d+`)

func parseLine(line string) ([]position, error) {
	const (
		matchNum = 2
	)

	res := make([]position, 0, matchNum)

	matches := reg.FindAllString(line, -1)
	if len(matches) != matchNum {
		return nil, errors.New("wrong coordinates line")
	}

	for i := range matches {
		coordinates, err := parseCoordinates(matches[i])
		if err != nil {
			return nil, fmt.Errorf("parse coordinates: %w", err)
		}

		res = append(res, coordinates)
	}

	return res, nil
}

func parseCoordinates(s string) (position, error) {
	const (
		cNum  = 2
		delim = ","
		xpos  = 0
		ypos  = 1
	)

	spl := strings.Split(s, delim)
	if len(spl) != cNum {
		return position{}, errors.New("wrong coordinates pair")
	}

	x, err := strconv.Atoi(spl[xpos])
	if err != nil {
		return position{}, fmt.Errorf("parse x to int: %w", err)
	}

	y, err := strconv.Atoi(spl[ypos])
	if err != nil {
		return position{}, fmt.Errorf("parse y to int: %w", err)
	}

	return position{
		x: x,
		y: y,
	}, nil
}

type filterFunc func(l line) bool

func part1Filter(l line) bool {
	return l.start.x == l.end.x || l.start.y == l.end.y
}

func filterLines(lines []line, filter filterFunc) []line {
	filtered := lines[:0]
	for _, x := range lines {
		if filter(x) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}
