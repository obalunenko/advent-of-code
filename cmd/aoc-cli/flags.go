package main

import (
	"github.com/urfave/cli/v2"
)

const (
	flagElapsed        = "elapsed"
	flagShortElapsed   = "e"
	flagBenchmark      = "bench"
	flagShortBenchmark = "b"
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

	res = append(res, &elapsed, &benchmark)

	return res
}
