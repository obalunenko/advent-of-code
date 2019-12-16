// Package day01 solves https://adventofcode.com/2019/day/1
package day01

import (
	"bufio"
	"io"
	"strconv"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

type solution struct {
	name string
}

func init() {
	puzzleName, err := puzzles.MakeName("2019", "day01")
	if err != nil {
		panic(err)
	}

	puzzles.Register(puzzleName, solution{
		name: puzzleName,
	})
}

func (s solution) Part1(input io.Reader) (string, error) {
	return calc(input, calcPart1)
}

func (s solution) Part2(input io.Reader) (string, error) {
	return calc(input, calcPart2)
}

func (s solution) Name() string {
	return s.name
}

type module struct {
	mass int
}

func (m module) fuel() int {
	mass := m.mass

	diff := mass % 3
	if diff != 0 {
		mass = mass - diff
	}

	f := (mass / 3) - 2

	return f
}

type calcFunc func(in chan module, res chan int, done chan struct{})

func calc(input io.Reader, calcFn calcFunc) (string, error) {
	var (
		lines int
		mass  int
		sum   int
		err   error
	)

	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calcFn(in, res, done)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		mass, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return "", err
		}

		in <- module{
			mass: mass,
		}
		lines++
	}

	close(in)

	if err = scanner.Err(); err != nil {
		return "", errors.Wrap(err, "scanner error")
	}

	for lines > 0 {
		select {
		case r := <-res:
			sum += r
		case <-done:
			lines--
		}
	}

	close(res)

	return strconv.Itoa(sum), nil
}

func calcPart1(in chan module, res chan int, done chan struct{}) {
	for i := range in {
		go func(m module, res chan int, done chan struct{}) {
			f := m.fuel()

			res <- f

			done <- struct{}{}
		}(i, res, done)
	}
}

func calcPart2(in chan module, res chan int, done chan struct{}) {
	for i := range in {
		go func(m module, res chan int, done chan struct{}) {
			var isDone bool

			for !isDone {
				f := m.fuel()

				res <- f

				if f/3 > 1 {
					m.mass = f
				} else {
					isDone = true
				}
			}

			done <- struct{}{}
		}(i, res, done)
	}
}
