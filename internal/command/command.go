// Package command holds functionality for running puzzle solving command.
package command

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/input"
)

// Run runs puzzle solving for passed year/day date.
func Run(ctx context.Context, year, day string) (puzzles.Result, error) {
	const timeout = time.Second * 30

	cli := input.NewFetcher(http.DefaultClient, timeout)

	return run(ctx, cli, year, day)
}

func run(ctx context.Context, cli input.Fetcher, year, day string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(year, day)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to get solver: %w", err)
	}

	fullName, err := puzzles.MakeName(s.Year(), s.Day())
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to make full name: %w", err)
	}

	asset, err := cli.Fetch(ctx, input.Date{
		Year: year,
		Day:  day,
	}, SessionFromContext(ctx))
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to get input for puzzle: %w", err)
	}

	opts := OptionsFromContext(ctx)

	res, err := puzzles.Solve(s, bytes.NewReader(asset), opts...)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to run [%s]: %w", fullName, err)
	}

	return res, nil
}
