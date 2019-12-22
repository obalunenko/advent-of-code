package day04

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
)

func init() {
	puzzleName, err := puzzles.MakeName("2019", "day04")
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
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(input); err != nil {
		return "", errors.Wrap(err, "failed to read")
	}

	limits := strings.Split(buf.String(), "-") // should be 2: low and high
	if len(limits) != 2 {
		return "", errors.New("invalid number of limits")
	}

	passwords, err := findPasswords(limits[0], limits[1])
	if err != nil {
		return "", errors.Wrap(err, "failed to find passwords")
	}

	return strconv.Itoa(passwords), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func (s solution) Name() string {
	return s.name
}

func findPasswords(low, high string) (int, error) {
	lowd, err := strconv.Atoi(low)
	if err != nil {
		return -1, errors.Wrap(err, "failed to convert low to int")
	}

	highd, err := strconv.Atoi(high)
	if err != nil {
		return -1, errors.Wrap(err, "failed to convert high to int")
	}

	pwds := make([]int, 0, highd-lowd)

	for i := lowd; i <= highd; i++ {
		if isPassword(i) {
			pwds = append(pwds, i)
		}
	}

	return len(pwds), nil
}

func isIncreasing(n int) bool {
	nmbs := intToSlice(n)

	for i := 1; i <= len(nmbs)-1; i++ {
		if nmbs[i] < nmbs[i-1] {
			return false
		}
	}

	return true
}

func hasDouble(n int) bool {
	nmbs := intToSlice(n)

	var hasDouble bool

	for i := 1; i <= len(nmbs)-1; i++ {
		if nmbs[i] == nmbs[i-1] {
			hasDouble = true
		}
	}

	return hasDouble
}

func isPassword(n int) bool {
	return isIncreasing(n) && hasDouble(n)
}

func intToSlice(n int) [6]int {
	return [6]int{
		(n % 1000000) / 100000,
		(n % 100000) / 10000,
		(n % 10000) / 1000,
		(n % 1000) / 100,
		(n % 100) / 10,
		n % 10,
	}
}
