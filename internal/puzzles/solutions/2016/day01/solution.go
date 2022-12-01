// Package day01 contains solution for https://adventofcode.com/2016/day/1 puzzle.
package day01

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

func (s solution) Year() string {
	return puzzles.Year2016.String()
}

func init() {
	puzzles.Register(solution{})
}

func (s solution) Part1(input io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(input); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	c := newCab()

	go func() {
		c.n.start()
	}()

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

	c.n.stop()

	l := c.n.Pos().manhattan()

	return strconv.Itoa(l), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(input); err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	c := newCab()

	go c.n.start()

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

	c.n.stop()

	rl := c.n.revisitedList()
	if len(rl) == 0 {
		return "", errors.New("no revisited points")
	}

	// get first
	l := rl[0].manhattan()

	return strconv.Itoa(l), nil
}

type turn string

const (
	leftTurn  = "L"
	rightTurn = "R"
)

var errInvalidTurn = errors.New("invalid turn value")

func turnFromstring(s string) (turn, error) {
	switch s {
	case leftTurn:
		return leftTurn, nil
	case rightTurn:
		return rightTurn, nil
	default:
		return "", errInvalidTurn
	}
}

type position struct {
	x, y int
}

func (p *position) addX(n int) {
	p.x += n
}

func (p *position) addY(n int) {
	p.y += n
}

func (p *position) subX(n int) {
	p.x -= n
}

func (p *position) subY(n int) {
	p.y -= n
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
	mu     *sync.Mutex
	curDir direction
	n      navigator
}

func newCab() cab {
	return cab{
		mu:     &sync.Mutex{},
		curDir: northDirection,
		n:      newNavigator(),
	}
}

var errInvalidDirect = errors.New("invalid direction")

const (
	step = 1
)

func (c *cab) Move(t turn, steps int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.curDir = c.curDir.strikeTo(t)
	if !c.curDir.isValid() {
		return errInvalidDirect
	}

	switch c.curDir {
	case northDirection:
		c.n.moveNorth(steps)
	case eastDirection:
		c.n.moveEast(steps)
	case southDirection:
		c.n.moveSouth(steps)
	case westDirection:
		c.n.moveWest(steps)
	case unknownDirection, sentinelDirection:
		return errInvalidDirect
	}

	return nil
}

type navigator struct {
	record    chan position
	pos       position
	track     track
	mu        *sync.Mutex
	wg        *sync.WaitGroup
	revisited []position
}

func newNavigator() navigator {
	return navigator{
		record: make(chan position),
		pos: position{
			x: 0,
			y: 0,
		},
		track:     newTrack(),
		mu:        &sync.Mutex{},
		wg:        &sync.WaitGroup{},
		revisited: []position{},
	}
}

func (n *navigator) recordTrack(p position) {
	n.mu.Lock()

	defer func() {
		n.mu.Unlock()
	}()

	if n.track.isVisited(p) {
		n.revisited = append(n.revisited, p)
	}

	n.track.record(p)
}

func (n *navigator) start() {
	n.wg.Add(1)

	for p := range n.record {
		n.recordTrack(p)
	}

	n.wg.Done()
}

func (n *navigator) stop() {
	close(n.record)

	n.wg.Wait()
}

func (n navigator) revisitedList() []position {
	return n.revisited
}

func (n *navigator) moveNorth(steps int) {
	for i := 0; i < steps; i++ {
		n.mu.Lock()
		n.pos.addY(step)
		n.mu.Unlock()

		n.record <- n.Pos()
	}
}

func (n *navigator) moveEast(steps int) {
	for i := 0; i < steps; i++ {
		n.mu.Lock()
		n.pos.addX(step)
		n.mu.Unlock()

		n.record <- n.Pos()
	}
}

func (n *navigator) moveSouth(steps int) {
	for i := 0; i < steps; i++ {
		n.mu.Lock()
		n.pos.subY(step)
		n.mu.Unlock()

		n.record <- n.Pos()
	}
}

func (n *navigator) moveWest(steps int) {
	for i := 0; i < steps; i++ {
		n.mu.Lock()
		n.pos.subX(step)
		n.mu.Unlock()

		n.record <- n.Pos()
	}
}

func (n navigator) Pos() position {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.pos
}

// Example: L4, R5
var re = regexp.MustCompile(`(?msi)([LR])(\d+)`)

const (
	_ = iota
	turnPos
	stepsPos

	totalMatchesNum = 3
)

var errInvalidCMD = errors.New("invalid command")

func splitCommand(cmd string) (turn, int, error) {
	parts := re.FindStringSubmatch(cmd)
	if len(parts) != totalMatchesNum {
		return "", 0, errInvalidCMD
	}

	t, err := turnFromstring(parts[turnPos])
	if err != nil {
		return "", 0, fmt.Errorf("turnFromstring: %w", err)
	}

	s, err := strconv.Atoi(parts[stepsPos])
	if err != nil {
		return "", 0, fmt.Errorf("invalid steps num: %w", err)
	}

	return t, s, nil
}

type track struct {
	t map[position]bool
	m *sync.Mutex
}

func newTrack() track {
	return track{
		t: make(map[position]bool),
		m: new(sync.Mutex),
	}
}

func (t track) record(p position) {
	t.m.Lock()
	defer t.m.Unlock()

	t.t[p] = true
}

func (t track) isVisited(p position) bool {
	t.m.Lock()
	defer t.m.Unlock()

	return t.t[p]
}
