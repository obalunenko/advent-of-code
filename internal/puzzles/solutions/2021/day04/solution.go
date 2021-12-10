// Package day04 contains solution for https://adventofcode.com/2021/day/4 puzzle.
package day04

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"sync"

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
	ctx := context.Background()

	game, err := newBingoGame(input)
	if err != nil {
		return "", fmt.Errorf("new bingo game: %w", err)
	}

	won, num := game.start(ctx, rule(1))

	res := won.sumMarked() * num

	return strconv.Itoa(res), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	ctx := context.Background()

	game, err := newBingoGame(input)
	if err != nil {
		return "", fmt.Errorf("new bingo game: %w", err)
	}

	won, num := game.start(ctx, rule(len(game.boards)))

	res := won.sumMarked() * num

	return strconv.Itoa(res), nil
}

func rule(boardsNum int) winRule {
	var count int

	return func(w winner) bool {
		count++

		return count == boardsNum
	}
}

type bingo struct {
	input  []int
	boards []*board
}

type winRule func(w winner) bool

func (b *bingo) start(ctx context.Context, rule winRule) (wonBoard *board, lastNum int) {
	players := make([]*player, 0, len(b.boards))

	in := make(chan int)
	boardWin := make(chan winner)

	ctx, cancel := context.WithCancel(ctx)

	for i := range b.boards {
		players = append(players, newPlayer(ctx, i, boardWin, b.boards[i]))
	}

	go func() {
		for _, n := range b.input {
			select {
			case <-ctx.Done():
				close(in)

				return
			default:
				in <- n
			}
		}

		close(in)
	}()

	go func() {
		for n := range in {
			for i := range players {
				p := players[i]

				p.input() <- n
			}
		}
	}()

	realWin := make(chan winner)

	go checkWinner(cancel, boardWin, realWin, rule)

	w := <-realWin

	wonBoard = b.boards[w.id]

	lastNum = w.num

	return wonBoard, lastNum
}

func checkWinner(cancelFunc context.CancelFunc, in, out chan winner, rule winRule) {
	for {
		w := <-in

		if rule(w) {
			cancelFunc()

			out <- w

			return
		}
	}
}

type player struct {
	mu     *sync.Mutex
	id     int
	in     chan int
	win    chan winSig
	active bool
	b      *board
}

func (p *player) isActive() bool {
	return p.active
}

func (p *player) setInactive() {
	p.active = false
}

func (p *player) input() chan int {
	return p.in
}

func (p *player) play(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			p.mu.Lock()

			p.setInactive()

			p.mu.Unlock()
		case num := <-p.in:
			p.mu.Lock()
			if !p.isActive() {
				p.mu.Unlock()

				continue
			}

			pos, ok := p.b.isPresent(num)
			if ok {
				p.b.state.update(ctx, pos)

				p.b.numbers[pos.vertical][pos.horizontal].setMarked()
			}

			if p.b.state.isWon() {
				p.win <- winSig{
					num: num,
				}

				p.setInactive()
			}

			p.mu.Unlock()
		}
	}
}

type winner struct {
	id  int
	num int
}

type winSig struct {
	num int
}

func newPlayer(ctx context.Context, id int, wonSig chan winner, b *board) *player {
	p := player{
		mu:     &sync.Mutex{},
		id:     id,
		in:     make(chan int),
		win:    make(chan winSig),
		active: true,
		b:      b,
	}

	go func() {
		sig := <-p.win

		wonSig <- winner{
			id:  p.id,
			num: sig.num,
		}
	}()

	go p.play(ctx)

	return &p
}

const (
	boardSize = 5
)

type number struct {
	val      int
	isMarked bool
}

func (n *number) setMarked() {
	n.isMarked = true
}

type board struct {
	id      int
	numbers [boardSize][boardSize]number
	state   state
}

func (b board) sumMarked() int {
	var sum int

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			num := b.numbers[i][j]
			if !num.isMarked {
				sum += num.val
			}
		}
	}

	return sum
}

type position struct {
	horizontal, vertical int
}

func (p position) String() string {
	return fmt.Sprintf("horizont=%d; vertical=%d", p.horizontal, p.vertical)
}

type state struct {
	verticals   [boardSize]int
	horizontals [boardSize]int
}

func newState() state {
	return state{
		verticals:   [boardSize]int{},
		horizontals: [boardSize]int{},
	}
}

func (s *state) String() string {
	return fmt.Sprintf("verticals=%v; horizontals=%v", s.verticals, s.horizontals)
}

func (s *state) update(_ context.Context, p position) {
	s.verticals[p.horizontal]++
	s.horizontals[p.vertical]++
}

func (s *state) isWon() bool {
	for i := 0; i < boardSize; i++ {
		if s.verticals[i] == boardSize || s.horizontals[i] == boardSize {
			return true
		}
	}

	return false
}

func (b board) isPresent(n int) (position, bool) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.numbers[i][j].val == n {
				return position{
					vertical:   i,
					horizontal: j,
				}, true
			}
		}
	}

	return position{}, false
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
			boardsNum++

			bg.boards = append(bg.boards, newBoard(boardsNum))

			cursor = 0
		case boardLine:
			for i, n := range numbers {
				bg.boards[boardsNum-1].numbers[cursor][i] = number{
					val:      n,
					isMarked: false,
				}
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

func newBoard(id int) *board {
	return &board{
		id:      id,
		numbers: [boardSize][boardSize]number{},
		state:   newState(),
	}
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
