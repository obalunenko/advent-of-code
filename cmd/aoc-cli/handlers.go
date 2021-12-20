package main

import (
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

	"github.com/obalunenko/advent-of-code/internal/command"
	"github.com/obalunenko/advent-of-code/internal/puzzles"
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
		ctx = command.ContextWithOptions(ctx, optionsFromCli(c)...)
		ctx = command.ContextWithSession(ctx, sessionFromCli(c))

		years := puzzles.GetYears()

		items := append(years, exit)

		searcher := func(input string, index int) bool {
			itm := items[index]

			itm = strings.ReplaceAll(strings.ToLower(itm), " ", "")

			input = strings.ReplaceAll(strings.ToLower(input), " ", "")

			return strings.Contains(itm, input)
		}

		prompt := promptui.Select{
			Label:             "Years menu (exit' for exit)",
			Items:             items,
			Size:              pageSize,
			CursorPos:         0,
			IsVimMode:         false,
			HideHelp:          false,
			HideSelected:      false,
			Templates:         nil,
			Keys:              nil,
			Searcher:          searcher,
			StartInSearchMode: false,
			Pointer:           promptui.DefaultCursor,
			Stdin:             nil,
			Stdout:            nil,
		}

		return handleYearChoices(ctx, prompt)
	}
}

func menuPuzzle(ctx context.Context, year string) error {
	solvers := puzzles.DaysByYear(year)

	items := append(solvers, back, exit)

	searcher := func(input string, index int) bool {
		itm := items[index]

		itm = strings.ReplaceAll(strings.ToLower(itm), " ", "")

		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(itm, input)
	}

	prompt := promptui.Select{
		Label:             "Puzzles menu (exit' for exit; back - to return to year selection)",
		Items:             items,
		Size:              pageSize,
		CursorPos:         0,
		IsVimMode:         false,
		HideHelp:          false,
		HideSelected:      false,
		Templates:         nil,
		Keys:              nil,
		Searcher:          searcher,
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

		res, err := command.Run(ctx, year, choice)
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

func sessionFromCli(c *cli.Context) string {
	var sess string

	sess = c.String(flagSession)
	if sess != "" {
		return sess
	}

	sess = c.String(flagShortSession)
	if sess != "" {
		return sess
	}

	return ""
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
