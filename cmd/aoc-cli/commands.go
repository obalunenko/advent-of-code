package main

import (
	"github.com/urfave/cli/v3"
)

func commands() []*cli.Command {
	const (
		cmdRun = "run"
	)

	cmds := []*cli.Command{
		{
			Name:   cmdRun,
			Usage:  "Runs advent-of-code application",
			Action: menu,
		},
	}

	return cmds
}
