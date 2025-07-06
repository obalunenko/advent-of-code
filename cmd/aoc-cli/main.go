// aoc-cli is a tool to run solutions to get answers for input on advent-of-code site.
package main

import (
	"context"
	"errors"
	"os"

	log "github.com/obalunenko/logger"
	"github.com/urfave/cli/v3"

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

	ctx = log.ContextWithLogger(ctx, log.FromContext(ctx))

	app := cli.Command{}
	app.Name = "aoc-cli"
	app.Description = "Solutions of puzzles for Advent Of Code (https://adventofcode.com/)\n" +
		"This command line tool contains solutions for puzzles and cli tool to run solutions to get " +
		"answers for input on site."
	app.Usage = `A command line tool for get solution for Advent of Code puzzles`
	app.Authors = []any{
		"Oleg Balunenko <oleg.balunenko@gmail.com>",
	}

	app.CommandNotFound = notFound
	app.Commands = commands()
	app.Version = printVersion(ctx)
	app.Before = printHeader
	app.After = onExit

	if err := app.Run(ctx, os.Args); err != nil {
		log.WithError(ctx, err).Fatal("Run failed")
	}
}
