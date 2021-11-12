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

	return getPassword(kpd, input)
}

func part1(input io.Reader) (string, error) {
	kpd := loadKeypadPart1()

	return getPassword(kpd, input)
}

func getPassword(kpd keypad, input io.Reader) (string, error) {
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

type keypadPos struct {
	x int
	y int
}

type move string

const (
	up    = move("U")
	down  = move("D")
	left  = move("L")
	right = move("R")
)

type num struct {
	val     string
	borders map[move]bool
}

type keypad struct {
	finger keypadPos
	specs  map[keypadPos]num
}

// TODO(obalunenko): refactor to load keypad from specs.

func loadKeypadPart2() keypad {
	/*

		keyboard

		       1
		     2 3 4
		   5 6 7 8 9
		     A B C
		       D

		start at `5`

		`7` is 0,0
	*/

	start := keypadPos{
		x: -2,
		y: 0,
	}

	instructions := map[keypadPos]num{
		{
			x: 0,
			y: 2,
		}: {
			val:     "1",
			borders: allowedMoves(down),
		},
		{
			x: -1,
			y: 1,
		}: {
			val:     "2",
			borders: allowedMoves(right, down),
		},
		{
			x: 0,
			y: 1,
		}: {
			val:     "3",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 1,
			y: 1,
		}: {
			val:     "4",
			borders: allowedMoves(left, down),
		},
		{
			x: -2,
			y: 0,
		}: {
			val:     "5",
			borders: allowedMoves(right),
		},
		{
			x: -1,
			y: 0,
		}: {
			val:     "6",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 0,
			y: 0,
		}: {
			val:     "7",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 1,
			y: 0,
		}: {
			val:     "8",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 2,
			y: 0,
		}: {
			val:     "9",
			borders: allowedMoves(left),
		},
		{
			x: -1,
			y: -1,
		}: {
			val:     "A",
			borders: allowedMoves(right, up),
		},
		{
			x: 0,
			y: -1,
		}: {
			val:     "B",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 1,
			y: -1,
		}: {
			val:     "C",
			borders: allowedMoves(left, up),
		},
		{
			x: 0,
			y: -2,
		}: {
			val:     "D",
			borders: allowedMoves(up),
		},
	}

	return keypad{
		finger: start,
		specs:  instructions,
	}
}

func loadKeypadPart1() keypad {
	/*
		keyboard
		    1 2 3
		    4 5 6
		    7 8 9

		let's predict that this is x y coordinates and 5 is a 0,0
	*/
	start := keypadPos{
		x: 0,
		y: 0,
	}

	instructions := map[keypadPos]num{
		{
			x: -1,
			y: 1,
		}: {
			val:     "1",
			borders: allowedMoves(right, down),
		},
		{
			x: 0,
			y: 1,
		}: {
			val:     "2",
			borders: allowedMoves(right, left, down),
		},
		{
			x: 1,
			y: 1,
		}: {
			val:     "3",
			borders: allowedMoves(left, down),
		},
		{
			x: -1,
			y: 0,
		}: {
			val:     "4",
			borders: allowedMoves(right, down, up),
		},
		{
			x: 0,
			y: 0,
		}: {
			val:     "5",
			borders: allowedMoves(right, left, down, up),
		},
		{
			x: 1,
			y: 0,
		}: {
			val:     "6",
			borders: allowedMoves(left, down, up),
		},
		{
			x: -1,
			y: -1,
		}: {
			val:     "7",
			borders: allowedMoves(right, up),
		},
		{
			x: 0,
			y: -1,
		}: {
			val:     "8",
			borders: allowedMoves(left, up, right),
		},
		{
			x: 1,
			y: -1,
		}: {
			val:     "9",
			borders: allowedMoves(left, up),
		},
	}

	return newKeypad(instructions, start)
}

func allowedMoves(am ...move) map[move]bool {
	allowed := map[move]bool{
		right: false,
		left:  false,
		down:  false,
		up:    false,
	}

	for _, m := range am {
		allowed[m] = true
	}

	return allowed
}

func newKeypad(specs map[keypadPos]num, startPos keypadPos) keypad {
	return keypad{
		finger: startPos,
		specs:  specs,
	}
}

func (k *keypad) move(m move) error {
	switch m {
	case up:
		if k.canMoveUp() {
			k.finger.y++
		}

	case down:
		if k.canMoveDown() {
			k.finger.y--
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
	return k.specs[k.finger].borders[right]
}

func (k keypad) canMoveLeft() bool {
	return k.specs[k.finger].borders[left]
}

func (k keypad) canMoveUp() bool {
	return k.specs[k.finger].borders[up]
}

func (k keypad) canMoveDown() bool {
	return k.specs[k.finger].borders[down]
}

func (k keypad) numb() string {
	return k.specs[k.finger].val
}
