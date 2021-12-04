// Package day03 contains solution for https://adventofcode.com/2021/day/3 puzzle.
package day03

import (
	"bufio"
	"fmt"
	"io"
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
	return puzzles.Day03.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	diagnostic, err := makeDiagnostic(input)
	if err != nil {
		return "", fmt.Errorf("make diagnostic report: %w", err)
	}

	r := findRates(diagnostic)

	consumption, err := r.consumption()
	if err != nil {
		return "", fmt.Errorf("consumption calculate: %w", err)
	}

	return consumption, nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

func makeDiagnostic(input io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(input)

	var report []string

	for scanner.Scan() {
		line := scanner.Text()

		report = append(report, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return report, nil
}

type bitrates struct {
	gamma   string
	epsilon string
}

func (b bitrates) consumption() (string, error) {
	g, err := strconv.ParseInt(b.gamma, 2, 64)
	if err != nil {
		return "", fmt.Errorf("parse gamma: %w", err)
	}

	e, err := strconv.ParseInt(b.epsilon, 2, 64)
	if err != nil {
		return "", fmt.Errorf("parse epsilon: %w", err)
	}

	c := g * e

	return strconv.FormatInt(c, 10), nil
}

func findRates(diagnostic []string) bitrates {
	var result bitrates

	type (
		bit string
		pos int
	)

	const (
		bit0 bit = "0"
		bit1 bit = "1"
	)

	bitsStat := make(map[pos]map[bit]int)

	for _, s := range diagnostic {
		for p, r := range []rune(s) {
			if bitsStat[pos(p)] == nil {
				bitsStat[pos(p)] = make(map[bit]int)
			}

			bitsStat[pos(p)][bit(r)]++
		}
	}

	for i := 0; i < len(bitsStat); i++ {
		m := bitsStat[pos(i)]

		if m[bit1] > m[bit0] {
			result.gamma += string(bit1)
			result.epsilon += string(bit0)
		} else {
			result.gamma += string(bit0)
			result.epsilon += string(bit1)
		}
	}

	return result
}
