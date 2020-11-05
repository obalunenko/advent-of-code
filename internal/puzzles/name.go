package puzzles

import (
	"errors"
	"os"
)

var (
	// ErrInvalidPzzlName means that such puzzle not exist.
	ErrInvalidPzzlName = errors.New("invalid puzzle name")
	// ErrInvalidYear means that such year not exist.
	ErrInvalidYear = errors.New("invalid year")
)

// MakeName builds puzzle name according to year and puzzle passed.
func MakeName(year string, puzzle string) (string, error) {
	if puzzle == "" {
		return "", ErrInvalidPzzlName
	}

	if year == "" {
		return "", ErrInvalidYear
	}

	return year + string(os.PathSeparator) + puzzle, nil
}
