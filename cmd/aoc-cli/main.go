// aoc-cli is a tool to run solutions to get answers for input on advent-of-code site.
package main

import (
	"context"
	"errors"
	"os"

	log "github.com/obalunenko/logger"
	"github.com/urfave/cli/v2"

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
	app.Authors = []*cli.Author{
		{
			Name:  "Oleg Balunenko",
			Email: "oleg.balunenko@gmail.com",
		},
	}
	app.CommandNotFound = notFound(ctx)
	app.Commands = commands(ctx)
	app.Version = printVersion(ctx)
	app.Before = printHeader(ctx)
	app.After = onExit(ctx)

	if err := app.Run(os.Args); err != nil {
		if errors.Is(err, errExit) {
			return
		}

		log.WithError(ctx, err).Fatal("Run failed")
	}
}
