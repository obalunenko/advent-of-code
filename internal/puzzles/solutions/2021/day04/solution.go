// Package day04 contains solution for https://adventofcode.com/2021/day/4 puzzle.
package day04

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
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
	return puzzles.Day04.String()
}

func (s solution) Part1(input io.Reader) (string, error) {
	_, err := newBingoGame(input)
	if err != nil {
		return "", fmt.Errorf("new bingo game: %w", err)
	}

	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(input io.Reader) (string, error) {
	return "", puzzles.ErrNotImplemented
}

type bingo struct {
	input  []int
	boards []*board
}

type board struct {
	numbers [5][5]int
}

type inputType int

const (
	_ inputType = iota
	inputNums
	boardLine
	emptyLine
)

func newBingoGame(input io.Reader) (*bingo, error) {
	scanner := bufio.NewScanner(input)

	var bg bingo

	var (
		idx       int
		boardsNum int
		cursor    int
	)

	for scanner.Scan() {
		it := boardLine

		if idx == 0 {
			it = inputNums
		}

		line := scanner.Text()
		if line == "" {
			it = emptyLine
		}

		numbers, err := getNumbers(line)
		if err != nil {
			return nil, fmt.Errorf("get numbers: %w", err)
		}

		switch it {
		case inputNums:
			bg.input = numbers
		case emptyLine:
			bg.boards = append(bg.boards, &board{
				numbers: [5][5]int{},
			})

			boardsNum++
			cursor = 0
		case boardLine:
			for i, n := range numbers {
				bg.boards[boardsNum-1].numbers[cursor][i] = n
			}

			cursor++
		}

		idx++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return &bg, nil
}

var re = regexp.MustCompile(`(?s)\d+`)

func getNumbers(s string) ([]int, error) {
	var nums []int

	for _, match := range re.FindAllString(s, -1) {
		n, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("parse num: %w", err)
		}

		nums = append(nums, n)
	}

	return nums, nil
}
