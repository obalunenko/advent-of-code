// Package declares common interface for puzzle solutions and functionality for register and run them.
package solutions

// Each solution should implement `Solver` interface, be implemented udnder separate package and contain `init()`
// function that will register that solution in list of all solvers.

// 	type solver struct {
//		name string
//	}
//
//	func init() {
// 		const puzzleName = "day01"
// 		puzzles.Register(puzzleName, solver{
//		name: puzzleName,
//		})
// 	}

// Then to register solution in the list of all solutions: make a blank import of package with puzzle solution
// at main.go
//
// _ "github.com/oleg-balunenko/advent-of-code/puzzles/solutions/day01"
