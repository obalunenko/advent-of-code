package main

import (
	"github.com/urfave/cli/v2"

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
		Name:        flagElapsed,
		Aliases:     []string{flagShortElapsed},
		Usage:       "Enables elapsed time metric",
		EnvVars:     nil,
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Value:       false,
		DefaultText: "",
		Destination: nil,
		HasBeenSet:  false,
	}

	benchmark := cli.BoolFlag{
		Name:        flagBenchmark,
		Aliases:     []string{flagShortBenchmark},
		Usage:       "Enables benchmark metric",
		EnvVars:     nil,
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Value:       false,
		DefaultText: "",
		Destination: nil,
		HasBeenSet:  false,
	}

	session := cli.StringFlag{
		Name:        flagSession,
		Aliases:     []string{flagShortSession},
		Usage:       "AOC auth session to get inputs",
		EnvVars:     []string{puzzles.AOCSession},
		FilePath:    "",
		Required:    true,
		Hidden:      false,
		TakesFile:   false,
		Value:       "",
		DefaultText: "",
		Destination: nil,
		HasBeenSet:  false,
	}

	res = append(res, &elapsed, &benchmark, &session)

	return res
}
