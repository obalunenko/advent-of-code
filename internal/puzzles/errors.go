package puzzles

import "errors"

var (
	// ErrInvalidPuzzleName means that such puzzle not exist.
	ErrInvalidPuzzleName = errors.New("invalid puzzle name")
	// ErrInvalidYear means that such year not exist.
	ErrInvalidYear = errors.New("invalid year")
	// ErrNotImplemented signal that puzzle in not implemented yet.
	ErrNotImplemented = errors.New("not implemented")
)
