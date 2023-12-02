package main

import (
	"context"

	"github.com/urfave/cli/v2"
)

func commands(ctx context.Context) []*cli.Command {
	const (
		cmdRun = "run"
		cmdGen = "gen"
	)

	cmds := []*cli.Command{
		{
			Name:                   cmdRun,
			Aliases:                nil,
			Usage:                  "Runs advent-of-code application",
			UsageText:              "",
			Description:            "",
			ArgsUsage:              "",
			Category:               "",
			BashComplete:           nil,
			Before:                 nil,
			After:                  nil,
			Action:                 menu(ctx),
			OnUsageError:           nil,
			Subcommands:            nil,
			Flags:                  cmdRunFlags(),
			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",
		},
		{
			Name:                   cmdGen,
			Aliases:                nil,
			Usage:                  "Generates new puzzle solution skeleton",
			UsageText:              "",
			Description:            "",
			ArgsUsage:              "",
			Category:               "",
			BashComplete:           nil,
			Before:                 nil,
			After:                  nil,
			Action:                 generate(ctx),
			OnUsageError:           nil,
			Subcommands:            nil,
			Flags:                  cmdGenFlags(),
			SkipFlagParsing:        false,
			HideHelp:               false,
			HideHelpCommand:        false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",
		},
	}

	return cmds
}
