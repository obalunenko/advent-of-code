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
	promptlist "github.com/manifoldco/promptui/list"
	log "github.com/obalunenko/logger"
	"github.com/savioxavier/termlink"
	"github.com/urfave/cli/v3"

	"github.com/obalunenko/advent-of-code/internal/command"
	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func onExit(_ context.Context, _ *cli.Command) error {
	fmt.Println("Exit...")

	return nil
}

func printHeader(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	const (
		padding  int  = 1
		minWidth int  = 0
		tabWidth int  = 0
		padChar  byte = ' '
	)

	w := tabwriter.NewWriter(cmd.Writer, minWidth, tabWidth, padding, padChar, tabwriter.TabIndent)

	_, err := fmt.Fprintf(w, `

 █████╗ ██████╗ ██╗   ██╗███████╗███╗   ██╗████████╗     ██████╗ ███████╗     ██████╗ ██████╗ ██████╗ ███████╗
██╔══██╗██╔══██╗██║   ██║██╔════╝████╗  ██║╚══██╔══╝    ██╔═══██╗██╔════╝    ██╔════╝██╔═══██╗██╔══██╗██╔════╝
███████║██║  ██║██║   ██║█████╗  ██╔██╗ ██║   ██║       ██║   ██║█████╗      ██║     ██║   ██║██║  ██║█████╗  
██╔══██║██║  ██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║       ██║   ██║██╔══╝      ██║     ██║   ██║██║  ██║██╔══╝  
██║  ██║██████╔╝ ╚████╔╝ ███████╗██║ ╚████║   ██║       ╚██████╔╝██║         ╚██████╗╚██████╔╝██████╔╝███████╗
╚═╝  ╚═╝╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝        ╚═════╝ ╚═╝          ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝

`)
	if err != nil {
		return ctx, fmt.Errorf("print version: %w", err)
	}

	return ctx, nil
}

func notFound(ctx context.Context, cmd *cli.Command, command string) {
	if _, err := fmt.Fprintf(
		cmd.Writer,
		"Command [%s] not supported.\nTry --help flag to see how to use it\n",
		command,
	); err != nil {
		log.WithError(ctx, err).Fatal("Failed to print not found message")
	}
}

func menu(ctx context.Context, c *cli.Command) error {
	ctx = command.ContextWithOptions(ctx, optionsFromCli(c)...)
	ctx = command.ContextWithSession(ctx, sessionFromCli(c))

	years := puzzles.GetYears()

	items := makeMenuItemsList(years, exit)

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
		Searcher:          searcher(items),
		StartInSearchMode: false,
		Pointer:           promptui.DefaultCursor,
		Stdin:             nil,
		Stdout:            nil,
	}

	return handleYearChoices(ctx, prompt)
}

func menuPuzzle(ctx context.Context, year string) error {
	solvers := puzzles.DaysByYear(year)

	items := makeMenuItemsList(solvers, back, exit)

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
		Searcher:          searcher(items),
		StartInSearchMode: false,
		Pointer:           promptui.DefaultCursor,
		Stdin:             nil,
		Stdout:            nil,
	}

	return handlePuzzleChoices(ctx, year, prompt)
}

func makeMenuItemsList(list []string, commands ...string) []string {
	items := make([]string, 0, len(list)+len(commands))

	items = append(items, list...)

	items = append(items, commands...)

	return items
}

func searcher(items []string) promptlist.Searcher {
	return func(input string, index int) bool {
		itm := items[index]

		itm = strings.ReplaceAll(strings.ToLower(itm), " ", "")

		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(itm, input)
	}
}

func handleYearChoices(ctx context.Context, opt promptui.Select) error {
	for {
		_, yearOpt, err := opt.Run()
		if err != nil {
			if isAbort(err) {
				return nil
			}

			return fmt.Errorf("prompt failed: %w", err)
		}

		if isExit(yearOpt) {
			return nil
		}

		err = menuPuzzle(ctx, yearOpt)
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
		_, dayOpt, err := opt.Run()
		if err != nil {
			if isAbort(err) {
				return errExit
			}

			return fmt.Errorf("prompt failed: %w", err)
		}

		if isExit(dayOpt) {
			return errExit
		}

		if isBack(dayOpt) {
			return nil
		}

		stopSpinner := setSpinner()

		res, err := command.Run(ctx, year, dayOpt)
		if err != nil {
			stopSpinner()

			if errors.Is(err, command.ErrUnauthorized) {
				fmt.Println(termlink.Link("Authorize here", "https://adventofcode.com/auth/login"))

				log.WithError(ctx, err).Fatal("Session expired")
			}

			log.WithError(ctx, err).Error("Puzzle run failed")

			continue
		}

		stopSpinner("Solved!")

		url := getURL(year, dayOpt)

		fmt.Println(res.String())

		fmt.Println(termlink.Link("Enter puzzle answers here", url))
	}
}

func getURL(year, day string) string {
	const urlFmt = "https://adventofcode.com/%s/day/%s"

	return fmt.Sprintf(urlFmt, year, day)
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

func optionsFromCli(c *cli.Command) []puzzles.RunOption {
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

func sessionFromCli(c *cli.Command) string {
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
func setSpinner() func(msg ...string) {
	const delayMs = 100

	s := spinner.New(
		spinner.CharSets[62],
		delayMs*time.Millisecond,
		spinner.WithFinalMSG(""),
		spinner.WithHiddenCursor(true),
		spinner.WithColor("yellow"),
		spinner.WithWriter(os.Stderr),
	)

	s.Prefix = "Solving in progress..."

	s.Start()

	return func(msg ...string) {
		if len(msg) != 0 {
			s.FinalMSG = msg[0]
		}

		s.Stop()
	}
}
