// Code generated by "stringer --type=Year --trimprefix=true --linecomment=true"; DO NOT EDIT.

package puzzles

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[yearUnknown-0]
	_ = x[Year2015-1]
	_ = x[Year2016-2]
	_ = x[Year2017-3]
	_ = x[Year2018-4]
	_ = x[Year2019-5]
	_ = x[Year2020-6]
	_ = x[Year2021-7]
	_ = x[Year2022-8]
	_ = x[Year2023-9]
	_ = x[yearSentinel-10]
}

const _Year_name = "yearUnknown201520162017201820192020202120222023yearSentinel"

var _Year_index = [...]uint8{0, 11, 15, 19, 23, 27, 31, 35, 39, 43, 47, 59}

func (i Year) String() string {
	if i < 0 || i >= Year(len(_Year_index)-1) {
		return "Year(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Year_name[_Year_index[i]:_Year_index[i+1]]
}
