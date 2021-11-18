// aoc-cli is a tool to run solutions to get answers for input on advent-of-code site.
package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	log "github.com/obalunenko/logger"
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
	app.Version = printVersion(ctx)
	app.Email = "oleg.balunenko@gmail.com"

	app.Flags = flags()

	app.Action = menu(ctx)
	app.Before = printHeader
	app.After = onExit

	if err := app.Run(os.Args); err != nil {
		if errors.Is(err, errExit) {
			return
		}

		log.WithError(ctx, err).Fatal("Run failed")
	}
}

const (
	flagElapsed        = "elapsed"
	flagShortElapsed   = "e"
	flagBenchmark      = "bench"
	flagShortBenchmark = "b"
)

func flags() []cli.Flag {
	var res []cli.Flag

	elapsed := cli.BoolFlag{
		Name:        fmt.Sprintf("%s, %s", flagElapsed, flagShortElapsed),
		Usage:       "Enables elapsed time metric",
		EnvVar:      "",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	benchmark := cli.BoolFlag{
		Name:        fmt.Sprintf("%s, %s", flagBenchmark, flagShortBenchmark),
		Usage:       "Enables benchmark metric",
		EnvVar:      "",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	res = append(res, elapsed, benchmark)

	return res
}

func onExit(_ *cli.Context) error {
	fmt.Println("Exit...")

	return nil
}

func menu(ctx context.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		ctx = contextWithOptions(ctx, optionsFromCli(c))

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

		stopSpinner := setSpinner()

		res, err := run(ctx, year, choice)
		if err != nil {
			log.WithError(ctx, err).Error("Puzzle run failed")

			stopSpinner()

			continue
		}

		stopSpinner()

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

func optionsFromCli(c *cli.Context) []puzzles.RunOption {
	const optsnum = 2

	options := make([]puzzles.RunOption, 0, optsnum)

	if c.GlobalBool(flagElapsed) || c.GlobalBool(flagShortElapsed) {
		options = append(options, puzzles.WithElapsed())
	}

	if c.GlobalBool(flagBenchmark) || c.GlobalBool(flagShortBenchmark) {
		options = append(options, puzzles.WithBenchmark())
	}

	return options
}

type optsCtxKey struct{}

func contextWithOptions(ctx context.Context, opts []puzzles.RunOption) context.Context {
	if len(opts) == 0 {
		return ctx
	}

	return context.WithValue(ctx, optsCtxKey{}, opts)
}

func optionsFromContext(ctx context.Context) []puzzles.RunOption {
	v := ctx.Value(optsCtxKey{})

	opts, ok := v.([]puzzles.RunOption)
	if !ok {
		return []puzzles.RunOption{}
	}

	return opts
}

func run(ctx context.Context, year, day string) (puzzles.Result, error) {
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

	opts := optionsFromContext(ctx)

	res, err := puzzles.Run(s, bytes.NewReader(asset), opts...)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to run [%s]: %w", fullName, err)
	}

	return res, nil
}

// setSpinner runs the displaying of spinner to handle long time operations. Returns stop func.
func setSpinner() func() {
	const delayms = 100

	s := spinner.New(
		spinner.CharSets[62],
		delayms*time.Millisecond,
		spinner.WithFinalMSG("Solved!"),
		spinner.WithHiddenCursor(true),
		spinner.WithColor("yellow"),
		spinner.WithWriter(os.Stderr),
	)

	s.Prefix = "Solving in progress..."

	s.Start()

	return func() {
		s.Stop()
	}
}
