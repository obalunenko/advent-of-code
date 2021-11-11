package day02

import (
	"bufio"
	"errors"
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
	return "", puzzles.ErrNotImplemented
}

func part1(input io.Reader) (string, error) {
	reader := bufio.NewReader(input)

	kpd := newKeypad()

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

			_, err = pwd.WriteString(strconv.Itoa(cur))
			if err != nil {
				return "", fmt.Errorf("write string: %w", err)
			}

			continue
		}

		if err = kpd.move(m); err != nil {
			return "", fmt.Errorf("move: %w", err)
		}
	}

	return pwd.String(), nil
}

/*
keyboard
    1 2 3
    4 5 6
    7 8 9

let's predict that this is x y coordinates and 5 is a 0,0
*/

type keypadPos struct {
	x int
	y int
}

type keypad struct {
	finger keypadPos
	dict   map[keypadPos]int
}

func newKeypad() keypad {
	dict := map[keypadPos]int{
		{
			x: -1,
			y: 1,
		}: 1,
		{
			x: 0,
			y: 1,
		}: 2,
		{
			x: 1,
			y: 1,
		}: 3,
		{
			x: -1,
			y: 0,
		}: 4,
		{
			x: 0,
			y: 0,
		}: 5,
		{
			x: 1,
			y: 0,
		}: 6,
		{
			x: -1,
			y: -1,
		}: 7,
		{
			x: 0,
			y: -1,
		}: 8,
		{
			x: 1,
			y: -1,
		}: 9,
	}

	return keypad{
		finger: keypadPos{
			x: 0,
			y: 0,
		},
		dict: dict,
	}

}

func (k *keypad) move(move string) error {
	const (
		up    = "U"
		down  = "D"
		left  = "L"
		right = "R"
	)

	switch move {
	case up:
		if k.canMoveYUp() {
			k.finger.y++
		}

	case down:
		if k.canMoveYDown() {
			k.finger.y--
		}

	case left:
		if k.canMoveXLeft() {
			k.finger.x--
		}
	case right:
		if k.canMoveXRight() {
			k.finger.x++
		}
	default:
		return fmt.Errorf("unsupported move")
	}

	return nil
}

func (k keypad) canMoveXRight() bool {
	return k.finger.x < 1
}

func (k keypad) canMoveXLeft() bool {
	return k.finger.x > -1
}

func (k keypad) canMoveYUp() bool {
	return k.finger.y < 1
}

func (k keypad) canMoveYDown() bool {
	return k.finger.y > -1
}

func (k keypad) numb() int {
	return k.dict[k.finger]
}
