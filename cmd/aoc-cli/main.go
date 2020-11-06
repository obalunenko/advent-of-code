package main

import (
	"bytes"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/advent-of-code/internal/input"
	"github.com/oleg-balunenko/advent-of-code/internal/puzzles"

	// register all solutions.
	_ "github.com/oleg-balunenko/advent-of-code/internal/puzzles/solutions"
)

const (
	exit = "exit"
	back = "back"
)

var (
	logLevel = flag.String("log_level", "INFO", "Set level of output logs")
)

func main() {
	defer func() {
		fmt.Println("Exiting...")
	}()

	flag.Parse()

	printVersion()

	setLogger()

	if err := menu(); err != nil {
		log.Fatal(err)
	}
}

func menu() error {
	years := puzzles.GetYears()

	prompt := promptui.Select{
		Label:             nil,
		Items:             append(years, exit),
		Size:              0,
		CursorPos:         0,
		IsVimMode:         false,
		HideHelp:          false,
		HideSelected:      false,
		Templates:         nil,
		Keys:              nil,
		Searcher:          nil,
		StartInSearchMode: false,
		Pointer:           nil,
		Stdin:             nil,
		Stdout:            nil,
	}

	return handleYearChoices(prompt)
}

func handleYearChoices(opt promptui.Select) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			return errors.Wrap(err, "prompt failed")
		}

		if isExit(choice) {
			return nil
		}

		err = menuPuzzle(choice)
		if errors.Is(err, errExit) {
			return nil
		}

		if err != nil {
			log.Error(err)

			continue
		}
	}
}

func setLogger() {
	l, err := log.ParseLevel(*logLevel)
	if err != nil {
		l = log.InfoLevel
	}

	log.SetLevel(l)

	formatter := &log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          true,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    true,
		PadLevelText:              false,
		QuoteEmptyFields:          true,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
	log.SetFormatter(formatter)
}

func menuPuzzle(year string) error {
	solvers := puzzles.NamesByYear(year)

	pageSize := 20

	prompt := promptui.Select{
		Label:             "Puzzles menu (exit' for exit; back - to return to year selection)",
		Items:             append(solvers, back, exit),
		Size:              pageSize,
		CursorPos:         0,
		IsVimMode:         false,
		HideHelp:          false,
		HideSelected:      false,
		Templates:         nil,
		Keys:              nil,
		Searcher:          nil,
		StartInSearchMode: false,
		Pointer:           promptui.DefaultCursor,
		Stdin:             nil,
		Stdout:            nil,
	}

	return handlePuzzleChoices(year, prompt)
}

func handlePuzzleChoices(year string, opt promptui.Select) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			return errors.Wrap(err, "prompt failed")
		}

		if isExit(choice) {
			return errExit
		}

		if isBack(choice) {
			return nil
		}

		res, err := run(year, choice)
		if err != nil {
			log.Error(err)

			continue
		}

		log.WithFields(log.Fields{
			"year":  res.Year,
			"name":  res.Name,
			"part1": res.Part1,
			"part2": res.Part2,
		}).Info("Puzzle answers")
	}
}

func isExit(input string) bool {
	return strings.EqualFold(exit, input)
}

func isBack(input string) bool {
	return strings.EqualFold(back, input)
}

var errExit = errors.New("exit is chosen")

func run(year string, name string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(year, name)
	if err != nil {
		return puzzles.Result{}, errors.Wrap(err, "failed to get solver")
	}

	fullName, err := puzzles.MakeName(s.Year(), s.Name())
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to make full name: %w", err)
	}

	asset, err := input.Asset(filepath.Clean(
		filepath.Join(input.InputDir, fmt.Sprintf("%s.txt", fullName))))
	if err != nil {
		return puzzles.Result{}, errors.Wrap(err, "failed to open input data")
	}

	res, err := puzzles.Run(s, bytes.NewReader(asset))
	if err != nil {
		return puzzles.Result{}, errors.Wrapf(err, "failed to run [%s]", s.Name())
	}

	return res, nil
}
