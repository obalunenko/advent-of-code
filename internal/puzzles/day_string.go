// Code generated by "stringer --type=Day --trimprefix=true --linecomment=true"; DO NOT EDIT.

package puzzles

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[dayUnknown-0]
	_ = x[Day01-1]
	_ = x[Day02-2]
	_ = x[Day03-3]
	_ = x[Day04-4]
	_ = x[Day05-5]
	_ = x[Day06-6]
	_ = x[Day07-7]
	_ = x[Day08-8]
	_ = x[Day09-9]
	_ = x[Day10-10]
	_ = x[Day11-11]
	_ = x[Day12-12]
	_ = x[Day13-13]
	_ = x[Day14-14]
	_ = x[Day15-15]
	_ = x[Day16-16]
	_ = x[Day17-17]
	_ = x[Day18-18]
	_ = x[Day19-19]
	_ = x[Day20-20]
	_ = x[Day21-21]
	_ = x[Day22-22]
	_ = x[Day23-23]
	_ = x[Day24-24]
	_ = x[Day25-25]
	_ = x[daySentinel-26]
}

const _Day_name = "dayUnknown12345678910111213141516171819202122232425daySentinel"

var _Day_index = [...]uint8{0, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 62}

func (i Day) String() string {
	if i < 0 || i >= Day(len(_Day_index)-1) {
		return "Day(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Day_name[_Day_index[i]:_Day_index[i+1]]
}
