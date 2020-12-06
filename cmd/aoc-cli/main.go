package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/obalunenko/advent-of-code/internal/input"
	"github.com/obalunenko/advent-of-code/internal/puzzles"
	// register all solutions.
	_ "github.com/obalunenko/advent-of-code/internal/puzzles/solutions"
)

const (
	exit     = "exit"
	back     = "back"
	pageSize = 10
	abort    = "^C"
)

var errExit = errors.New("exit is chosen")

func main() {
	defer func() {
		fmt.Println("Exiting...")
	}()

	app := cli.NewApp()
	app.Name = "aoc-cli"

	app.Description = "Solutions of puzzles for Advent Of Code (https://adventofcode.com/)\n" +
		"This command line tool contains solutions for puzzles and cli tool to run solutions to get " +
		"answers for input on site."
	app.Usage = `a command line tool for get solution for Advent of Code puzzles`
	app.Author = "Oleg Balunenko"
	app.Version = versionInfo()
	app.Email = "oleg.balunenko@gmail.com"
	app.Flags = globalFlags()
	app.Action = menu

	if err := app.Run(os.Args); err != nil {
		if errors.Is(err, errExit) {
			return
		}

		log.Fatal(err)
	}
}

func globalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "log_level",
			Usage:       "Level of output logs",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       log.InfoLevel.String(),
			Destination: nil,
		},
	}
}

func setLogger(ctx *cli.Context) {
	formatter := log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          true,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}

	log.SetFormatter(&formatter)

	lvl, err := log.ParseLevel(ctx.GlobalString("log_level"))
	if err != nil {
		lvl = log.InfoLevel
	}

	log.SetLevel(lvl)
}

func menu(ctx *cli.Context) error {
	setLogger(ctx)

	years := puzzles.GetYears()

	prompt := promptui.Select{
		Label:             "Years menu (exit' for exit)",
		Items:             append(years, exit),
		Size:              pageSize,
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
			if isAbort(err) {
				return nil
			}

			return errors.Wrap(err, "prompt failed")
		}

		if isExit(choice) {
			return nil
		}

		err = menuPuzzle(choice)
		if err != nil {
			if errors.Is(err, errExit) {
				return nil
			}

			log.Error(err)

			continue
		}
	}
}

func menuPuzzle(year string) error {
	solvers := puzzles.NamesByYear(year)

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
			if isAbort(err) {
				return errExit
			}

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

func isAbort(err error) bool {
	return strings.HasSuffix(err.Error(), abort)
}

func isBack(input string) bool {
	return strings.EqualFold(back, input)
}

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
