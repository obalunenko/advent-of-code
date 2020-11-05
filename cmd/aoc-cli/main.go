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

func menu() error {
	solvers := puzzles.Solvers()

	pageSize := 20

	prompt := promptui.Select{
		Label:             "Puzzles menu (input 'exit' for exit)",
		Items:             append(solvers, exit),
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

	return handleChoices(prompt)
}

func handleChoices(opt promptui.Select) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			return errors.Wrap(err, "prompt failed")
		}

		if isExit(choice) {
			return nil
		}

		res, err := run(choice)
		if err != nil {
			log.Error(err)

			continue
		}

		log.WithFields(log.Fields{
			"name":  res.Name,
			"part1": res.Part1,
			"part2": res.Part2,
		}).Info("Puzzle answers")
	}
}

func isExit(input string) bool {
	return strings.EqualFold(exit, input)
}

func run(puzzle string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(puzzle)
	if err != nil {
		return puzzles.Result{}, errors.Wrap(err, "failed to get solver")
	}

	asset, err := input.Asset(filepath.Clean(
		filepath.Join(input.InputDir, fmt.Sprintf("%s.txt", s.Name()))))
	if err != nil {
		return puzzles.Result{}, errors.Wrap(err, "failed to open input data")
	}

	res, err := puzzles.Run(s, bytes.NewReader(asset))
	if err != nil {
		return puzzles.Result{}, errors.Wrapf(err, "failed to run [%s]", s.Name())
	}

	return res, nil
}
