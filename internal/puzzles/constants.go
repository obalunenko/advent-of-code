package puzzles

//go:generate stringer --type=Day --trimprefix=true --linecomment=true

// Day presents here for purpose of documentation.
type Day int

const (
	dayUnknown Day = iota

	Day01 // 1
	Day02 // 2
	Day03 // 3
	Day04 // 4
	Day05 // 5
	Day06 // 6
	Day07 // 7
	Day08 // 8
	Day09 // 9
	Day10 // 10
	Day11 // 11
	Day12 // 12
	Day13 // 13
	Day14 // 14
	Day15 // 15
	Day16 // 16
	Day17 // 17
	Day18 // 18
	Day19 // 19
	Day20 // 20
	Day21 // 21
	Day22 // 22
	Day23 // 23
	Day24 // 24
	Day25 // 25

	daySentinel
)

//go:generate stringer --type=Year --trimprefix=true --linecomment=true

// Year presents here for purpose of documentation.
type Year int

const (
	yearUnknown Year = iota

	Year2015 // 2015
	Year2016 // 2016
	Year2017 // 2017
	Year2018 // 2018
	Year2019 // 2019
	Year2020 // 2020
	Year2021 // 2021
	Year2022 // 2022
	Year2023 // 2023

	yearSentinel
)

const (
	unknown    = "unknown"
	unsolved   = "not solved"
	undefined  = "undefined"
	inProgress = "in progress"

	// AOCSession env variable name.
	AOCSession = "AOC_SESSION"
)
