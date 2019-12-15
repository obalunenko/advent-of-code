package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/puzzles"
	_ "github.com/oleg-balunenko/advent-of-code/puzzles/solutions/day01"
)

var (
	inputDir = flag.String(
		"input_dir",
		filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "oleg-balunenko", "advent-of-code", "input"),
		"Path to directory with puzzles input files")
)

func main() {
	flag.Parse()

	menu()
}

func menu() {
	fmt.Println("Menu:")

	solvers := puzzles.Solvers()

	var choices = make(map[string]string, len(solvers))

	for i, s := range solvers {
		choices[strconv.Itoa(i+1)] = s
		fmt.Printf("%d. %s \n", i+1, s)
	}

	var text string

	for {
		if _, err := fmt.Scanln(&text); err != nil {
			log.Fatal(err)
		}

		if puz, ok := choices[text]; !ok {
			fmt.Println("wrong choice, try again")
		} else {
			run(puz)
			return
		}
	}
}

func run(name string) {
	s, err := puzzles.GetSolver(name)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get solver"))
	}

	input := filepath.Join(*inputDir, fmt.Sprintf("%s.txt", name))

	if err := puzzles.Run(s, input); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run puzzle solver"))
	}
}
