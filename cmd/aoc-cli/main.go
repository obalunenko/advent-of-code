// aoc-cli is a tool to run solutions to get answers for input on advent-of-code site.
package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/obalunenko/logger"
	"github.com/obalunenko/version"
	"github.com/urfave/cli"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/input"
	_ "github.com/obalunenko/advent-of-code/internal/puzzles/solutions" // register all solutions.
)

const (
	exit     = "exit"
	back     = "back"
	pageSize = 10
	abort    = "^C"
)

var errExit = errors.New("exit is chosen")

func main() {
	ctx := context.Background()

	app := cli.NewApp()
	app.Name = "aoc-cli"

	app.Description = "Solutions of puzzles for Advent Of Code (https://adventofcode.com/)\n" +
		"This command line tool contains solutions for puzzles and cli tool to run solutions to get " +
		"answers for input on site."
	app.Usage = `a command line tool for get solution for Advent of Code puzzles`
	app.Author = "Oleg Balunenko"
	app.Version = version.GetVersion()
	app.Email = "oleg.balunenko@gmail.com"

	app.Action = menu(ctx)
	app.Before = printVersion
	app.After = onExit

	if err := app.Run(os.Args); err != nil {
		if errors.Is(err, errExit) {
			return
		}

		log.WithError(ctx, err).Fatal("Run failed")
	}
}

func onExit(_ *cli.Context) error {
	fmt.Println("Exit...")

	return nil
}

func menu(ctx context.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
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

		return handleYearChoices(ctx, prompt)
	}
}

func handleYearChoices(ctx context.Context, opt promptui.Select) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			if isAbort(err) {
				return nil
			}

			return fmt.Errorf("prompt failed: %w", err)
		}

		if isExit(choice) {
			return nil
		}

		err = menuPuzzle(ctx, choice)
		if err != nil {
			if errors.Is(err, errExit) {
				return nil
			}

			log.WithError(ctx, err).Error("Puzzle menu failed")

			continue
		}
	}
}

func menuPuzzle(ctx context.Context, year string) error {
	solvers := puzzles.DaysByYear(year)

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

	return handlePuzzleChoices(ctx, year, prompt)
}

func handlePuzzleChoices(ctx context.Context, year string, opt promptui.Select) error {
	for {
		_, choice, err := opt.Run()
		if err != nil {
			if isAbort(err) {
				return errExit
			}

			return fmt.Errorf("prompt failed: %w", err)
		}

		if isExit(choice) {
			return errExit
		}

		if isBack(choice) {
			return nil
		}

		res, err := run(year, choice)
		if err != nil {
			log.WithError(ctx, err).Error("Puzzle run failed")

			continue
		}

		fmt.Println(res.String())
	}
}

func isExit(in string) bool {
	return strings.EqualFold(exit, in)
}

func isAbort(err error) bool {
	return strings.HasSuffix(err.Error(), abort)
}

func isBack(in string) bool {
	return strings.EqualFold(back, in)
}

func run(year, day string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(year, day)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to get solver: %w", err)
	}

	fullName, err := puzzles.MakeName(s.Year(), s.Day())
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to make full name: %w", err)
	}

	asset, err := input.Asset(fmt.Sprintf("%s.txt", fullName))
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to open input data: %w", err)
	}

	res, err := puzzles.Run(s, bytes.NewReader(asset))
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to run [%s]: %w", fullName, err)
	}

	return res, nil
}
