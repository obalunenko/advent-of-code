// Package day02 contains solution for https://adventofcode.com/2016/day/2 puzzle.
package day02

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	return puzzles.Year2016.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	return part1(input)
}

func (s solution) Part2(input io.Reader) (string, error) {
	return part2(input)
}

func part2(input io.Reader) (string, error) {
	kpd := loadKeypadPart2()

	return getDoorCode(kpd, input)
}

func part1(input io.Reader) (string, error) {
	kpd := loadKeypadPart1()

	return getDoorCode(kpd, input)
}

func getDoorCode(kpd keypad, input io.Reader) (string, error) {
	reader := bufio.NewReader(input)

	var pwd strings.Builder

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
			cur := kpd.numb()

			_, err = pwd.WriteString(cur)
			if err != nil {
				return "", fmt.Errorf("write string: %w", err)
			}

			continue
		}

		if err = kpd.move(move(m)); err != nil {
			return "", fmt.Errorf("move: %w", err)
		}
	}

	return pwd.String(), nil
}

type move string

const (
	up    = move("U")
	down  = move("D")
	left  = move("L")
	right = move("R")
)

type keypadPos struct {
	x int
	y int
}

type grid [][]string

type keypad struct {
	finger keypadPos
	grid   grid
}

/*
loadKeypadPart2
keyboard

	    1
	  2 3 4
	5 6 7 8 9
	  A B C
	    D

start at `5`

let's predict that this is a 2 dimension matrix and '5' is 0,2m
*/
func loadKeypadPart2() keypad {
	start := keypadPos{
		x: 0,
		y: 2,
	}

	g := [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	return keypad{
		finger: start,
		grid:   g,
	}
}

/*
loadKeypadPart1
keyboard

	1 2 3
	4 5 6
	7 8 9

let's predict that this is a 2 dimension matrix and 5 is a 1,1
*/
func loadKeypadPart1() keypad {
	g := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	start := keypadPos{
		x: 1,
		y: 1,
	}

	return newKeypad(g, start)
}

func newKeypad(specs grid, startPos keypadPos) keypad {
	return keypad{
		finger: startPos,
		grid:   specs,
	}
}

func (k *keypad) move(m move) error {
	switch m {
	case up:
		if k.canMoveUp() {
			k.finger.y--
		}

	case down:
		if k.canMoveDown() {
			k.finger.y++
		}

	case left:
		if k.canMoveLeft() {
			k.finger.x--
		}
	case right:
		if k.canMoveRight() {
			k.finger.x++
		}
	default:
		return fmt.Errorf("unsupported move")
	}

	return nil
}

func (k keypad) canMoveRight() bool {
	cur := k.finger

	if cur.x == len(k.grid[cur.y])-1 {
		return false
	}

	return k.grid[cur.y][cur.x+1] != ""
}

func (k keypad) canMoveLeft() bool {
	cur := k.finger

	if cur.x == 0 {
		return false
	}

	return k.grid[cur.y][cur.x-1] != ""
}

func (k keypad) canMoveUp() bool {
	cur := k.finger

	if cur.y == 0 {
		return false
	}

	return k.grid[cur.y-1][cur.x] != ""
}

func (k keypad) canMoveDown() bool {
	cur := k.finger

	if cur.y == len(k.grid)-1 {
		return false
	}

	return k.grid[cur.y+1][cur.x] != ""
}

func (k keypad) numb() string {
	return k.grid[k.finger.y][k.finger.x]
}
