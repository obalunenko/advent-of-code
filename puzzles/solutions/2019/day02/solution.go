package day02

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"

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

	c.replace(map[int]int{
		1: 12,
		2: 2,
	})

	res, err := c.calc()
	if err != nil {
		return "", errors.Wrap(err, "failed to calc")
	}

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}

type computer struct {
	input map[int]int
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
	c.input = make(map[int]int, len(nums))

	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return c, errors.Wrap(err, "failed to convert string to int")
		}

		c.input[i] = n
	}

	return c, nil
}

func (c computer) calc() (int, error) {
	var (
		result int
		err    error
	)

loop:
	for i := 0; i < len(c.input); i += shift {
		opt, aPos, bPos, resPos := c.input[i], c.input[i+1], c.input[i+2], c.input[i+3]
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
			return -1, errors.Errorf("not supported opt code [%d] at pos [%d]", opt, i)
		}
	}

	return result, err
}

func (c computer) add(aPos, bPos, resPos int) error {
	a, ok := c.input[aPos]
	if !ok {
		return errors.New("value not exist")
	}

	b, ok := c.input[bPos]
	if !ok {
		return errors.New("value not exist")
	}

	res := a + b
	c.input[resPos] = res

	return nil
}

func (c *computer) mult(aPos, bPos, resPos int) error {
	a, ok := c.input[aPos]
	if !ok {
		return errors.New("value not exist")
	}

	b, ok := c.input[bPos]
	if !ok {
		return errors.New("value not exist")
	}

	res := a * b
	c.input[resPos] = res

	return nil
}

func (c *computer) abort() (int, error) {
	res, ok := c.input[0]
	if !ok {
		return 0, errors.New("value not exist")
	}

	return res, nil
}

func (c *computer) replace(data map[int]int) {
	for i, v := range data {
		c.input[i] = v
	}
}
