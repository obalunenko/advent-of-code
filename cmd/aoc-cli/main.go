package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
	_ "github.com/oleg-balunenko/advent-of-code/puzzles/solutions/day01"
)

const (
	exit = "exit"
)

func main() {
	defer func() {
		fmt.Println("Exiting...")
	}()

	printVersion()

	flag.Parse()

	if err := menu(); err != nil {
		log.Fatal(err)
	}
}

func menu() error {
	path, err := inputPath()
	if err != nil {
		return errors.Wrap(err, "failed to get puzzles input directory")
	}

	if isExit(path) {
		return nil
	}

	solvers := puzzles.Solvers()

	prompt := promptui.Select{
		Label:             "Puzzles menu (input 'exit' for exit)",
		Items:             append(solvers, exit),
		Size:              0,
		IsVimMode:         false,
		HideHelp:          false,
		HideSelected:      false,
		Templates:         nil,
		Keys:              nil,
		Searcher:          nil,
		StartInSearchMode: false,
		Pointer:           promptui.DefaultCursor,
		Stdin:             nil,
		Stdout:            nil,
	}

	for {
		_, result, err := prompt.Run()
		if err != nil {
			return errors.Wrap(err, "prompt failed")
		}

		if isExit(result) {
			return nil
		}

		if err := run(result, path); err != nil {
			return errors.Wrap(err, "failed run puzzle")
		}
	}
}

func isExit(input string) bool {
	return strings.EqualFold(exit, input)
}

func inputPath() (string, error) {
	validate := func(input string) error {
		if isExit(input) {
			return nil
		}

		if _, err := os.Stat(input); os.IsNotExist(err) {
			return errors.Wrap(err, "invalid input directory path")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Puzzles input path (input 'exit' for exit)",
		Default:   "",
		AllowEdit: false,
		Validate:  validate,
		Mask:      0,
		Templates: nil,
		IsConfirm: false,
		IsVimMode: false,
		Pointer:   promptui.DefaultCursor,
	}

	path, err := prompt.Run()
	if err != nil {
		return "", errors.Wrap(err, "path prompt failed")
	}

	return path, nil
}

func run(puzzle string, inputdir string) error {
	s, err := puzzles.GetSolver(puzzle)
	if err != nil {
		return errors.Wrap(err, "failed to get solver")
	}

	input := filepath.Clean(
		filepath.Join(inputdir, fmt.Sprintf("%s.txt", puzzle)),
	)

	if err := puzzles.Run(s, input); err != nil {
		return errors.Wrap(err, "failed to run puzzle solver")
	}

	return nil
}
