package day01

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day01"
	year       = "2016"
)

type solution struct {
	year string
	name string
}

func (s solution) Name() string {
	return s.name
}

func (s solution) Year() string {
	return s.year
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

func (s solution) Part1(input io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(input); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	c := newCab()

	cmds := strings.Split(buf.String(), ", ")
	for _, cmd := range cmds {
		t, s, err := splitCommand(cmd)
		if err != nil {
			return "", fmt.Errorf("split command: %w", err)
		}

		if err = c.Move(t, s); err != nil {
			return "", fmt.Errorf("move: %w", err)
		}
	}

	l := c.Pos().manhattan()

	return strconv.Itoa(l), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

type turn string

const (
	leftTurn  = "L"
	rightTurn = "R"
)

func tunrFromstring(s string) (turn, error) {
	switch s {
	case leftTurn:
		return leftTurn, nil
	case rightTurn:
		return rightTurn, nil
	default:
		return "", errors.New("invalid turn value")
	}
}

type position struct {
	x, y int
}

func (p *position) addX(n int) {
	p.x = p.x + n
}

func (p *position) addY(n int) {
	p.y = p.y + n
}

func (p *position) subX(n int) {
	p.x = p.x - n
}

func (p *position) subY(n int) {
	p.y = p.y - n
}

func (p position) manhattan() int {
	x, y := p.x, p.y

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}

type direction uint

const (
	unknownDirection direction = iota

	northDirection
	eastDirection
	southDirection
	westDirection

	sentinelDirection
)

func (d direction) isValid() bool {
	return d > unknownDirection && d < sentinelDirection
}

func (d direction) strikeTo(t turn) direction {
	switch t {
	case rightTurn:
		if d == westDirection {
			return northDirection
		}

		return d + 1
	case leftTurn:
		if d == northDirection {
			return westDirection
		}

		return d - 1

	default:
		return unknownDirection
	}
}

type cab struct {
	pos    position
	curDir direction
}

func newCab() cab {
	return cab{
		pos: position{
			x: 0,
			y: 0,
		},
		curDir: northDirection,
	}
}
func (c *cab) Move(t turn, steps int) error {
	c.curDir = c.curDir.strikeTo(t)
	if !c.curDir.isValid() {
		return errors.New("invalid direction")
	}

	switch c.curDir {
	case northDirection:
		c.pos.addY(steps)
	case eastDirection:
		c.pos.addX(steps)
	case southDirection:
		c.pos.subY(steps)
	case westDirection:
		c.pos.subX(steps)
	}

	return nil
}

func (c cab) Pos() position {
	return c.pos
}

// Example: L4, R5
var re = regexp.MustCompile(`(?msi)(L|R)(\d+)`)

const (
	fullMatchPos = iota
	turnPos
	stepsPos

	totalMatchesNum = 3
)

func splitCommand(cmd string) (turn, int, error) {
	parts := re.FindStringSubmatch(cmd)
	if len(parts) != totalMatchesNum {
		return "", 0, errors.New("invalid command")
	}

	t, err := tunrFromstring(parts[turnPos])
	if err != nil {
		return "", 0, errors.New("invalid turn")
	}
	s, err := strconv.Atoi(parts[stepsPos])
	if err != nil {
		return "", 0, errors.New("invalid steps num")
	}

	return t, s, nil
}
