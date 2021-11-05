// Package day01 solves https://adventofcode.com/2019/day/1
package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

type solution struct{}

func (s solution) Year() string {
	return puzzles.Year2019.String()
}

func init() {
	puzzles.Register(solution{})
}

func (s solution) Part1(input io.Reader) (string, error) {
	return calc(input, calcPart1)
}

func (s solution) Part2(input io.Reader) (string, error) {
	return calc(input, calcPart2)
}

func (s solution) Day() string {
	return puzzles.Day01.String()
}

const (
	divFactor = 3
	subFactor = 2
)

type module struct {
	mass int
}

func (m module) fuel() int {
	mass := m.mass

	diff := mass % divFactor
	if diff != 0 {
		mass -= diff
	}

	f := (mass / divFactor) - subFactor

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
			return "", fmt.Errorf("faied to parse int: %w", err)
		}

		in <- module{
			mass: mass,
		}
		lines++
	}

	close(in)

	if err = scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
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
	const endNum = 1

	for i := range in {
		go func(m module, res chan int, done chan struct{}) {
			for {
				f := m.fuel()

				res <- f

				if f/divFactor <= endNum {
					break
				}

				m.mass = f
			}

			done <- struct{}{}
		}(i, res, done)
	}
}
