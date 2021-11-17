package puzzles

//go:generate stringer --type=Day --trimprefix=true --linecomment=true

// Day presents here for purpose of documentation.
type Day int

const (
	dayUnknown Day = iota

	Day01 // day01
	Day02 // day02
	Day03 // day03
	Day04 // day04
	Day05 // day05
	Day06 // day06
	Day07 // day07
	Day08 // day08
	Day09 // day09
	Day10 // day10
	Day11 // day11
	Day12 // day12
	Day13 // day13
	Day14 // day14
	Day15 // day15
	Day16 // day16
	Day17 // day17
	Day18 // day18
	Day19 // day19
	Day20 // day20
	Day21 // day21
	Day22 // day22
	Day23 // day23
	Day24 // day24
	Day25 // day25

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

	yearSentinel
)

const (
	unknown    = "unknown"
	unsolved   = "not solved"
	undefined  = "undefined"
	inProgress = "in progress"
)
