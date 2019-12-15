package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"

	"github.com/oleg-balunenko/advent-of-code/internal/puzzle"
	_ "github.com/oleg-balunenko/advent-of-code/internal/solutions/day01"
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

	solvers := puzzle.Solvers()

	var choices = make(map[string]string, len(solvers))

	for i, s := range solvers {
		choices[strconv.Itoa(i+1)] = s
		fmt.Printf("%d. %s \n", i+1, s)
	}

	var text string

	for {
		_, err := fmt.Scanln(&text)
		if err != nil {
			log.Fatal(err)
		}

		if pname, ok := choices[text]; !ok {
			fmt.Println("wrong choice, try again")
		} else {
			run(pname)
			return
		}
	}
}

func run(name string) {
	s, err := puzzle.GetSolver(name)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get solver"))
	}

	input := filepath.Join(*inputDir, fmt.Sprintf("%s.txt", name))

	if err := puzzle.Run(s, input); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run puzzle solver"))
	}
}
