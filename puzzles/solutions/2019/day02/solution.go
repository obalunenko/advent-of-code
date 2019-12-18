// Package day02 solves https://adventofcode.com/2019/day/2
package day02

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

func init() {
	puzzleName, err := puzzles.MakeName("2019", "day02")
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
	c, err := newComputer(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to init computer")
	}

	c.input(12, 2)

	res, err := c.calc()
	if err != nil {
		return "", errors.Wrap(err, "failed to calc")
	}

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	c, err := newComputer(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to init computer")
	}

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			c.reset()

			c.input(i, j)

			res, err := c.calc()
			if err != nil {
				return "", errors.Wrap(err, "failed to calc")
			}

			if res == 19690720 {
				log.WithFields(log.Fields{
					"noun": i,
					"verb": j,
				}).Info("Solved at positions")
				return strconv.Itoa(nounVerb(i, j)), nil
			}
		}
	}

	return "", errors.New("can't found non and verb")
}

func nounVerb(noun int, verb int) int {
	return 100*noun + verb
}

func (s solution) Name() string {
	return s.name
}

type computer struct {
	memory  map[int]int
	initial []int
}

const (
	optAdd   = 1
	optMult  = 2
	optAbort = 99

	shift = 4
)

func newComputer(input io.Reader) (computer, error) {
	var c computer

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(input); err != nil {
		return c, errors.Wrap(err, "failed to read")
	}

	nums := strings.Split(buf.String(), ",")
	c.initial = make([]int, len(nums))
	c.memory = make(map[int]int, len(nums))

	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return c, errors.Wrap(err, "failed to convert string to int")
		}

		c.initial[i] = n
		c.memory[i] = n
	}

	return c, nil
}

func (c computer) calc() (int, error) {
	var (
		result int
		err    error
	)

loop:
	for i := 0; i < len(c.memory); i += shift {
		opt, aPos, bPos, resPos := c.memory[i], c.memory[i+1], c.memory[i+2], c.memory[i+3]
		switch opt {
		case optAdd:
			if err = c.add(aPos, bPos, resPos); err != nil {
				return 0, errors.Wrap(err, "failed to add")
			}
		case optMult:
			if err = c.mult(aPos, bPos, resPos); err != nil {
				return 0, errors.Wrap(err, "failed to mult")
			}
		case optAbort:
			result, err = c.abort()
			break loop
		default:
			result = -1
			err = errors.Errorf("not supported opt code [%d] at pos [%d]", opt, i)
			break loop
		}
	}

	return result, err
}

func (c computer) add(aPos, bPos, resPos int) error {
	a, ok := c.memory[aPos]
	if !ok {
		return errors.New("value not exist")
	}

	b, ok := c.memory[bPos]
	if !ok {
		return errors.New("value not exist")
	}

	res := a + b
	c.memory[resPos] = res

	return nil
}

func (c *computer) mult(aPos, bPos, resPos int) error {
	a, ok := c.memory[aPos]
	if !ok {
		return errors.New("value not exist")
	}

	b, ok := c.memory[bPos]
	if !ok {
		return errors.New("value not exist")
	}

	res := a * b
	c.memory[resPos] = res

	return nil
}

func (c *computer) abort() (int, error) {
	res, ok := c.memory[0]
	if !ok {
		return 0, errors.New("value not exist")
	}

	return res, nil
}

func (c *computer) input(noun int, verb int) {
	c.memory[1] = noun
	c.memory[2] = verb
}

func (c *computer) reset() {
	c.memory = make(map[int]int, len(c.initial))

	for i, n := range c.initial {
		n := n
		c.memory[i] = n
	}
}
