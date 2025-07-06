package main

import (
	"github.com/urfave/cli/v3"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	flagElapsed        = "elapsed"
	flagShortElapsed   = "e"
	flagBenchmark      = "bench"
	flagShortBenchmark = "b"
	flagSession        = "session"
	flagShortSession   = "s"
)

func cmdRunFlags() []cli.Flag {
	var res []cli.Flag

	elapsed := cli.BoolFlag{
		Name:    flagElapsed,
		Usage:   "Enables elapsed time metric",
		Aliases: []string{flagShortElapsed},
	}

	benchmark := cli.BoolFlag{
		Name:    flagBenchmark,
		Usage:   "Enables benchmark metric",
		Aliases: []string{flagShortBenchmark},
	}

	session := cli.StringFlag{
		Name:     flagSession,
		Usage:    "AOC auth session to get inputs",
		Sources:  cli.EnvVars(puzzles.AOCSession),
		Required: true,
		Aliases:  []string{flagShortSession},
	}

	res = append(res, &elapsed, &benchmark, &session)

	return res
}
