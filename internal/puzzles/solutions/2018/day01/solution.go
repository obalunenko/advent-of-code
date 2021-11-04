package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day01"
	year       = "2018"
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

func (s solution) Part1(in io.Reader) (string, error) {
	return part1(in)
}

func (s solution) Part2(in io.Reader) (string, error) {
	return part2(in)
}

var (
	re = regexp.MustCompile(`(?s)(?P<sign>[+-])(?P<digits>\d+)`)
)

const (
	_ = iota
	sign
	digits

	totalmatches = 3
)

func part1(in io.Reader) (string, error) {
	scanner := bufio.NewScanner(in)

	var curfreq int

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)

		if len(matches) != totalmatches {
			return "", fmt.Errorf("wrong matches[%d] for line[%s], should be [%d]",
				len(matches), line, totalmatches)
		}

		d, err := strconv.Atoi(matches[digits])
		if err != nil {
			return "", fmt.Errorf("strconv atoi: %w", err)
		}

		switch matches[sign] {
		case "+":
			curfreq += d
		case "-":
			curfreq -= d
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	return strconv.Itoa(curfreq), nil
}

func part2(in io.Reader) (string, error) {
	var buf bytes.Buffer

	if _, err := buf.ReadFrom(in); err != nil {
		return "", fmt.Errorf("failed to read: %w", err)
	}

	b := buf.Bytes()

	seenfreqs := make(map[int]bool)

	var (
		curfreq int
		loops   int
		found   bool
	)

	for !found {
		if len(seenfreqs) == 0 {
			seenfreqs[curfreq] = true
		}

		scanner := bufio.NewScanner(bytes.NewReader(b))

		for scanner.Scan() {
			line := scanner.Text()

			matches := re.FindStringSubmatch(line)

			if len(matches) != totalmatches {
				return "", fmt.Errorf("wrong matches[%d] for line[%s], should be [%d]",
					len(matches), line, totalmatches)
			}

			d, err := strconv.Atoi(matches[digits])
			if err != nil {
				return "", fmt.Errorf("strconv atoi: %w", err)
			}

			switch matches[sign] {
			case "+":
				curfreq += d
			case "-":
				curfreq -= d
			}

			if seenfreqs[curfreq] {
				found = true

				break
			}

			seenfreqs[curfreq] = true
		}

		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("scanner error: %w", err)
		}

		loops++
	}

	return strconv.Itoa(curfreq), nil
}
