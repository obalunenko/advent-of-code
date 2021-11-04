package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"sync"

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

	fdevice := newDevice()

	for scanner.Scan() {
		line := scanner.Text()

		delta, err := getFreqDelta(line)
		if err != nil {
			return "", err
		}

		fdevice.Apply(delta)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	return strconv.Itoa(fdevice.CurFreq()), nil
}

func part2(in io.Reader) (string, error) {
	var buf bytes.Buffer

	if _, err := buf.ReadFrom(in); err != nil {
		return "", fmt.Errorf("failed to read: %w", err)
	}

	b := buf.Bytes()

	var (
		loops int
		found bool
	)

	fdevice := newDevice()

	for !found {
		scanner := bufio.NewScanner(bytes.NewReader(b))

		for scanner.Scan() {
			line := scanner.Text()

			delta, err := getFreqDelta(line)
			if err != nil {
				return "", fmt.Errorf("get frequency delta: %w", err)
			}

			fdevice.Apply(delta)

			if fdevice.SeenFreq(fdevice.CurFreq()) {
				found = true

				break
			}
		}

		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("scanner error: %w", err)
		}

		loops++
	}

	return strconv.Itoa(fdevice.CurFreq()), nil
}

type device struct {
	frequency int
	mu        sync.Mutex
	seen      map[int]int
}

func newDevice() *device {
	d := device{
		frequency: 0,
		mu:        sync.Mutex{},
		seen:      make(map[int]int),
	}

	d.seen[0] = 1

	return &d
}

func (d *device) Apply(delta freqDelta) {
	switch delta.sign {
	case "+":
		d.frequency += delta.d
	case "-":
		d.frequency -= delta.d
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.seen[d.frequency] = d.seen[d.frequency] + 1
}

func (d *device) CurFreq() int {
	return d.frequency
}

func (d *device) SeenFreq(freq int) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.seen[freq] > 1
}

type freqDelta struct {
	sign string
	d    int
}

func getFreqDelta(line string) (freqDelta, error) {
	matches := re.FindStringSubmatch(line)

	if len(matches) != totalmatches {
		return freqDelta{}, fmt.Errorf("wrong matches[%d] for line[%s], should be [%d]",
			len(matches), line, totalmatches)
	}

	d, err := strconv.Atoi(matches[digits])
	if err != nil {
		return freqDelta{}, fmt.Errorf("strconv atoi: %w", err)
	}

	return freqDelta{
		sign: matches[sign],
		d:    d,
	}, nil
}
