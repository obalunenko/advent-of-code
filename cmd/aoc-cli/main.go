package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
	// register all solutions
	_ "github.com/oleg-balunenko/advent-of-code/puzzles/solutions"
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
		EnvironmentOverrideColors: false,
		DisableTimestamp:          true,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    true,
		QuoteEmptyFields:          true,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
	log.SetFormatter(formatter)
}

func menu() error {
	path, err := inputPath()
	if err != nil {
		return errors.Wrap(err, "failed to get puzzles input directory")
	}

	if isExit(path) {
		return nil
	}

	solvers := puzzles.Solvers()

	prompt := promptui.Select{
		Label:             "Puzzles menu (input 'exit' for exit)",
		Items:             append(solvers, exit),
		Size:              20,
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

	return handleChoices(prompt, path)
}

func handleChoices(opt promptui.Select, inputDir string) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			return errors.Wrap(err, "prompt failed")
		}

		if isExit(choice) {
			return nil
		}

		res, err := run(choice, inputDir)
		if err != nil {
			log.Error(err)
			continue
		}

		log.WithFields(log.Fields{
			"name":  res.Puzzle,
			"part1": res.Part1,
			"part2": res.Part2,
		}).Info("Puzzle answers")
	}
}

func isExit(input string) bool {
	return strings.EqualFold(exit, input)
}

func inputPath() (string, error) {
	validate := func(input string) error {
		if isExit(input) {
			return nil
		}

		if _, err := os.Stat(input); os.IsNotExist(err) {
			return errors.Wrap(err, "invalid input directory path")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Puzzles input path (input 'exit' for exit)",
		Default:   "",
		AllowEdit: false,
		Validate:  validate,
		Mask:      0,
		Templates: nil,
		IsConfirm: false,
		IsVimMode: false,
		Pointer:   promptui.DefaultCursor,
	}

	path, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "path prompt failed")
	}

	return path, nil
}

func run(puzzle string, inputdir string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(puzzle)
	if err != nil {
		return puzzles.Result{}, errors.Wrap(err, "failed to get solver")
	}

	input := filepath.Clean(
		filepath.Join(inputdir, fmt.Sprintf("%s.txt", s.Name())),
	)

	res, err := puzzles.Run(s, input)
	if err != nil {
		return puzzles.Result{}, errors.Wrapf(err, "failed to run [%s]", s.Name())
	}

	return res, nil
}
