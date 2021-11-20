package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	log "github.com/obalunenko/logger"
	"github.com/urfave/cli/v2"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/input"
)

func onExit(_ context.Context) cli.AfterFunc {
	return func(c *cli.Context) error {
		fmt.Println("Exit...")

		return nil
	}
}

func printHeader(_ context.Context) cli.BeforeFunc {
	const (
		padding  int  = 1
		minWidth int  = 0
		tabWidth int  = 0
		padChar  byte = ' '
	)

	return func(c *cli.Context) error {
		w := tabwriter.NewWriter(c.App.Writer, minWidth, tabWidth, padding, padChar, tabwriter.TabIndent)

		_, err := fmt.Fprintf(w, `

 █████╗ ██████╗ ██╗   ██╗███████╗███╗   ██╗████████╗     ██████╗ ███████╗     ██████╗ ██████╗ ██████╗ ███████╗
██╔══██╗██╔══██╗██║   ██║██╔════╝████╗  ██║╚══██╔══╝    ██╔═══██╗██╔════╝    ██╔════╝██╔═══██╗██╔══██╗██╔════╝
███████║██║  ██║██║   ██║█████╗  ██╔██╗ ██║   ██║       ██║   ██║█████╗      ██║     ██║   ██║██║  ██║█████╗  
██╔══██║██║  ██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║       ██║   ██║██╔══╝      ██║     ██║   ██║██║  ██║██╔══╝  
██║  ██║██████╔╝ ╚████╔╝ ███████╗██║ ╚████║   ██║       ╚██████╔╝██║         ╚██████╗╚██████╔╝██████╔╝███████╗
╚═╝  ╚═╝╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝        ╚═════╝ ╚═╝          ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝

`)
		if err != nil {
			return fmt.Errorf("print version: %w", err)
		}

		return nil
	}
}

func notFound(ctx context.Context) cli.CommandNotFoundFunc {
	return func(c *cli.Context, command string) {
		if _, err := fmt.Fprintf(
			c.App.Writer,
			"Command [%s] not supported.\nTry --help flag to see how to use it\n",
			command,
		); err != nil {
			log.WithError(ctx, err).Fatal("Failed to print not found message")
		}
	}
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
	const optsNum = 2

	options := make([]puzzles.RunOption, 0, optsNum)

	if c.Bool(flagElapsed) || c.Bool(flagShortElapsed) {
		options = append(options, puzzles.WithElapsed())
	}

	if c.Bool(flagBenchmark) || c.Bool(flagShortBenchmark) {
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
	const delayMs = 100

	s := spinner.New(
		spinner.CharSets[62],
		delayMs*time.Millisecond,
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
