// Package day02 contains solution for https://adventofcode.com/2021/day/2 puzzle.
package day02

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"

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
	return puzzles.Day02.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	subm := newSubmarine()

	for scanner.Scan() {
		line := scanner.Text()

		act, err := parseAction(line)
		if err != nil {
			return "", fmt.Errorf("parse action: %w", err)
		}

		if err = subm.move(act); err != nil {
			return "", fmt.Errorf("submarine move: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	res := subm.pos.x * subm.pos.y

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

type action struct {
	move  move
	steps int
}

var moveRe = regexp.MustCompile(`(?s)(\w+)\s(\d+)`)

const (
	_ = iota
	movePos
	moveNumPos

	totalMatches = 3
)

var errInvalidFormat = errors.New("invalid action format")

func parseAction(s string) (action, error) {
	matches := moveRe.FindStringSubmatch(s)
	if len(matches) != totalMatches {
		return action{}, fmt.Errorf("[%s]: %w", s, errInvalidFormat)
	}

	m, err := parseMove(matches[movePos])
	if err != nil {
		return action{}, fmt.Errorf("parse move: %w", err)
	}

	n, err := strconv.Atoi(matches[moveNumPos])
	if err != nil {
		return action{}, fmt.Errorf("parse steps num: %w", err)
	}

	return action{
		move:  m,
		steps: n,
	}, nil
}

type position struct {
	x int // horizontal position
	y int // depth
}

type move string

const (
	moveUp      = "up"
	moveDown    = "down"
	moveForward = "forward"
)

func parseMove(s string) (move, error) {
	var m move

	switch s {
	case moveUp:
		m = moveUp
	case moveForward:
		m = moveForward
	case moveDown:
		m = moveDown
	default:
		return "", fmt.Errorf("[%s]: %w", s, errInvalidMove)
	}

	return m, nil
}

type submarine struct {
	pos position
}

func newSubmarine() submarine {
	return submarine{
		pos: position{
			x: 0,
			y: 0,
		},
	}
}

var errInvalidMove = errors.New("invalid move")

func (s *submarine) move(act action) error {
	switch act.move {
	case moveUp:
		s.pos.y -= act.steps
	case moveForward:
		s.pos.x += act.steps
	case moveDown:
		s.pos.y += act.steps
	default:
		return errInvalidMove
	}

	return nil
}
