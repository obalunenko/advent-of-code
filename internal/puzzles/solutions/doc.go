// Package solutions registers solutions of puzzles.
//
// Each solution should implement `Solver` interface, be implemented under separate package and contain `init()`
// function that will register that solution in list of all solvers.
//
// Example:
//
// 	type solver struct {
//		name string
//	}
//
//	func init() {
// 		puzzleName, err := solutions.MakeName("2019", "day01")
//		if err != nil {
//			panic(err)
//		}
// 		puzzles.Register(puzzleName, solver{
//		name: puzzleName,
//		})
// 	}
//
// Then to register solution in the list of all solutions: make a blank import of package with puzzle solution
// at register.go
//
// 	import  _ "github.com/obalunenko/advent-of-code/puzzles/solutions/day01"
//
// And then blank import solutions package at main.go to register all solutions
//
//  import _ "github.com/obalunenko/advent-of-code/puzzles/solutions
package solutions
