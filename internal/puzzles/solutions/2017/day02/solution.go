// Package day02 contains solution for https://adventofcode.com/2017/day/2 puzzle.
package day02

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	log "github.com/obalunenko/logger"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func init() {
	puzzles.Register(solution{})
}

type solution struct{}

func (s solution) Day() string {
	return puzzles.Day02.String()
}

func (s solution) Year() string {
	return puzzles.Year2017.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var checksum int

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.ReplaceAll(line, "\t", " ")

		var min, max int

		numbers := strings.Split(line, " ")

		for i, number := range numbers {
			d, err := strconv.Atoi(number)
			if err != nil {
				return "", fmt.Errorf("atoi: %w", err)
			}

			if i == 0 {
				min, max = d, d
			}

			log.WithFields(context.TODO(), log.Fields{
				"checksum": checksum,
				"min":      min,
				"max":      max,
				"i":        i,
				"number":   d,
				"row":      line,
			}).Info("Iterate")

			if d < min {
				min = d
			}

			if d > max {
				max = d
			}
		}

		checksum += max - min

	}
	return strconv.Itoa(checksum), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}
