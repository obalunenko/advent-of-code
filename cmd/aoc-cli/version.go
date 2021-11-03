// aoc-cli is a tool to run solutions to get answers for input on advent-of-code site.
package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/obalunenko/version"
	"github.com/urfave/cli"
)

func printVersion(_ *cli.Context) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	_, err := fmt.Fprintf(w, `
| app_name:	%s	|
| version:	%s	|
| short_commit:	%s	|
| commit:	%s	|
| build_date:	%s	|
| goversion:	%s	|
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

`,
		version.GetAppName(),
		version.GetVersion(),
		version.GetShortCommit(),
		version.GetCommit(),
		version.GetBuildDate(),
		version.GetGoVersion())
	if err != nil {
		return fmt.Errorf("print version: %w", err)
	}

	return nil
}
