package solutions

import (
	"errors"
	"os"
)

// MakeName builds puzzle name according to year and puzzle passed
func MakeName(year string, puzzle string) (string, error) {
	if puzzle == "" {
		return "", errors.New("invalid puzzle name")
	}

	if year == "" {
		return "", errors.New("invalid year")
	}

	return year + string(os.PathSeparator) + puzzle, nil
}
