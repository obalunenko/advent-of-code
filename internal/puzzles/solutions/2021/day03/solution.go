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

	r := findPowerConsumptionRates(diagnostic)

	consumption, err := r.consumption()
	if err != nil {
		return "", fmt.Errorf("consumption calculate: %w", err)
	}

	return consumption, nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	diagnostic, err := makeDiagnostic(input)
	if err != nil {
		return "", fmt.Errorf("make diagnostic report: %w", err)
	}

	r := lifeSupportRate(diagnostic)

	consumption, err := r.consumption()
	if err != nil {
		return "", fmt.Errorf("consumption calculate: %w", err)
	}

	return consumption, nil
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
	first  string
	second string
}

func (b bitrates) consumption() (string, error) {
	const (
		baseBin = 2
		baseDec = 10
		bitsize = 64
	)

	g, err := strconv.ParseInt(b.first, baseBin, bitsize)
	if err != nil {
		return "", fmt.Errorf("parse first: %w", err)
	}

	e, err := strconv.ParseInt(b.second, baseBin, bitsize)
	if err != nil {
		return "", fmt.Errorf("parse second: %w", err)
	}

	c := g * e

	return strconv.FormatInt(c, baseDec), nil
}

type (
	bit string
	pos int
)

const (
	bit0 bit = "0"
	bit1 bit = "1"
)

func findPowerConsumptionRates(diagnostic []string) bitrates {
	var result bitrates

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
			result.first += string(bit1)
			result.second += string(bit0)
		} else {
			result.first += string(bit0)
			result.second += string(bit1)
		}
	}

	return result
}

func lifeSupportRate(diagnostic []string) bitrates {
	o2 := lifeRate(diagnostic, o2Criteria)
	co2 := lifeRate(diagnostic, co2Criteria)

	return bitrates{
		first:  o2,
		second: co2,
	}
}

func lifeRate(diagnostic []string, criteriaFunc bitCriteriaFunc) string {
	var (
		result string
		idx    pos
	)

	for len(diagnostic) != 1 {
		bitstat := make(map[bit]int)

		for i := range diagnostic {
			runes := []rune(diagnostic[i])

			b := bit(runes[idx])

			bitstat[b]++
		}

		diagnostic = filterDiagnostic(diagnostic, idx, bitstat, criteriaFunc)

		idx++
	}

	result = diagnostic[0]

	return result
}

type bitCriteriaFunc func(bitstat map[bit]int) bit

func o2Criteria(bitstat map[bit]int) bit {
	common := bit1

	if bitstat[bit0] > bitstat[bit1] {
		common = bit0
	}

	return common
}

func co2Criteria(bitstat map[bit]int) bit {
	common := bit0

	if bitstat[bit1] < bitstat[bit0] {
		common = bit1
	}

	return common
}

func filterDiagnostic(diagnostic []string, idx pos, bitstat map[bit]int, bfunc bitCriteriaFunc) []string {
	common := bfunc(bitstat)

	filtered := make([]string, 0, len(diagnostic))

	for i := range diagnostic {
		runes := []rune(diagnostic[i])

		if bit(runes[idx]) == common {
			filtered = append(filtered, diagnostic[i])
		}
	}

	return filtered
}
