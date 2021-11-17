package puzzles

import (
	"path/filepath"
)

// MakeName builds puzzle name according to year and puzzle passed.
func MakeName(year, puzzle string) (string, error) {
	if puzzle == "" {
		return "", ErrInvalidPuzzleName
	}

	if year == "" {
		return "", ErrInvalidYear
	}

	return filepath.Join(year, puzzle), nil
}
