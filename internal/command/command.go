// Package command holds functionality for running puzzle solving command.
package command

import (
	"bytes"
	"context"
	"fmt"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
	"github.com/obalunenko/advent-of-code/internal/puzzles/input"
)

// Run runs puzzle solving for passed year/day date.
func Run(ctx context.Context, year, day string) (puzzles.Result, error) {
	s, err := puzzles.GetSolver(year, day)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to get solver: %w", err)
	}

	fullName, err := puzzles.MakeName(s.Year(), s.Day())
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to make full name: %w", err)
	}

	asset, err := input.Get(ctx, input.Date{
		Year: year,
		Day:  day,
	}, SessionFromContext(ctx))
	if err != nil {
		return puzzles.Result{}, err
	}

	opts := OptionsFromContext(ctx)

	res, err := puzzles.Solve(s, bytes.NewReader(asset), opts...)
	if err != nil {
		return puzzles.Result{}, fmt.Errorf("failed to run [%s]: %w", fullName, err)
	}

	return res, nil
}
