// Package day02 solves https://adventofcode.com/2019/day/2
package day02

import (
	"io"
	"strconv"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
	"github.com/oleg-balunenko/advent-of-code/puzzles/utils/intcomputer"
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
	c, err := intcomputer.New(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to init computer")
	}

	c.Input(12, 2)

	res, err := c.Execute()
	if err != nil {
		return "", errors.Wrap(err, "failed to calc")
	}

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	c, err := intcomputer.New(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to init computer")
	}

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			c.Reset()

			c.Input(i, j)

			res, err := c.Execute()
			if err != nil {
				return "", errors.Wrap(err, "failed to calc")
			}

			if res == 19690720 {
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
