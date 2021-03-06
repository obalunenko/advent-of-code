package day02

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day02"
	year       = "2020"
)

type solution struct {
	year string
	name string
}

func (s solution) Year() string {
	return s.year
}

func (s solution) Name() string {
	return s.name
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

var (
	pwdRegex = regexp.MustCompile(`(?s)(\d{1,2})-(\d{1,2}) ([a-zA-Z]): ([[:word:]]+)`)
)

func (s solution) Part1(input io.Reader) (string, error) {
	const (
		_ int = iota // full match - not needed
		matchMin
		matchMax
		matchChar
		matchPwd

		totalmatches = 5
	)

	type passwordParams struct {
		min  int
		max  int
		char string
	}

	type inparams struct {
		pwd       string
		pwdParams passwordParams
	}

	var count int

	scanner := bufio.NewScanner(input)

	inchan := make(chan inparams)
	reschan := make(chan bool)
	donechan := make(chan struct{})

	go func(in chan inparams, res chan bool, done chan struct{}) {
		for d := range in {
			go func(in inparams, res chan bool, done chan struct{}) {
				count := strings.Count(in.pwd, in.pwdParams.char)

				res <- count >= in.pwdParams.min && count <= in.pwdParams.max

				done <- struct{}{}
			}(d, res, done)
		}
	}(inchan, reschan, donechan)

	var operations int

	for scanner.Scan() {
		line := scanner.Text()
		submatch := pwdRegex.FindStringSubmatch(line)

		if len(submatch) != totalmatches {
			return "", fmt.Errorf("wrong matches[%d] for line[%s], should be [%d]",
				len(submatch), line, totalmatches)
		}

		pwd := submatch[matchPwd]

		min, err := strconv.Atoi(submatch[matchMin])
		if err != nil {
			return "", fmt.Errorf("failed to parse min[%s]: %w", submatch[matchMin], err)
		}

		max, err := strconv.Atoi(submatch[matchMax])
		if err != nil {
			return "", fmt.Errorf("failed to parse max[%s]: %w", submatch[matchMax], err)
		}

		params := passwordParams{
			min:  min,
			max:  max,
			char: submatch[matchChar],
		}

		in := inparams{
			pwd:       pwd,
			pwdParams: params,
		}

		inchan <- in

		operations++
	}

	close(inchan)

loop:
	for {
		select {
		case isMatch := <-reschan:
			if isMatch {
				count++

			}
		case <-donechan:
			operations--
		}
		if operations == 0 {
			break loop
		}
	}

	close(reschan)
	close(donechan)

	return strconv.Itoa(count), nil
}

func (s solution) Part2(input io.Reader) (string, error) {
	const (
		_ int = iota // full match - not needed
		matchFirstPos
		matchSecondPos
		matchChar
		matchPwd

		totalmatches = 5
	)

	type passwordParams struct {
		firstPos  int
		secondPos int
		char      string
	}

	type inparams struct {
		pwd       string
		pwdParams passwordParams
	}

	var count int

	scanner := bufio.NewScanner(input)

	inchan := make(chan inparams)
	reschan := make(chan bool)
	donechan := make(chan struct{})

	go func(in chan inparams, res chan bool, done chan struct{}) {
		for d := range in {
			go func(in inparams, res chan bool, done chan struct{}) {
				var found int
				if in.pwdParams.char == string([]rune(in.pwd)[in.pwdParams.firstPos-1]) {
					found++
				}

				if in.pwdParams.char == string([]rune(in.pwd)[in.pwdParams.secondPos-1]) {
					found++
				}

				res <- found > 0 && found < 2

				done <- struct{}{}
			}(d, res, done)
		}
	}(inchan, reschan, donechan)

	var operations int

	for scanner.Scan() {
		line := scanner.Text()
		submatch := pwdRegex.FindStringSubmatch(line)

		if len(submatch) != totalmatches {
			return "", fmt.Errorf("wrong matches[%d] for line[%s], should be [%d]",
				len(submatch), line, totalmatches)
		}

		pwd := submatch[matchPwd]

		firstPos, err := strconv.Atoi(submatch[matchFirstPos])
		if err != nil {
			return "", fmt.Errorf("failed to parse first pos[%s]: %w", submatch[matchFirstPos], err)
		}

		secondPos, err := strconv.Atoi(submatch[matchSecondPos])
		if err != nil {
			return "", fmt.Errorf("failed to parse second pos[%s]: %w", submatch[matchSecondPos], err)
		}

		params := passwordParams{
			firstPos:  firstPos,
			secondPos: secondPos,
			char:      submatch[matchChar],
		}

		in := inparams{
			pwd:       pwd,
			pwdParams: params,
		}

		inchan <- in

		operations++
	}

	close(inchan)

loop:
	for {
		select {
		case isMatch := <-reschan:
			if isMatch {
				count++

			}
		case <-donechan:
			operations--
		}
		if operations == 0 {
			break loop
		}
	}

	close(reschan)
	close(donechan)

	return strconv.Itoa(count), nil
}
