// Package getenv provides a simple way to get environment variables.
// It's type-safe and supports built-in types and slices of them.
// Note: empty environment variable values are treated as "not set".
//
// Slices are split by "," and time.Time values are parsed with the time.RFC3339 layout,
// unless WithSeparator or WithTimeLayout sets another one.
// Slice elements are not trimmed, so "1, 2" is not a valid []int,
// and a trailing separator adds an empty element: "a,b," is []string{"a", "b", ""}.
//
// The supported types are listed in Parsable.
package getenv

import (
	"errors"
	"fmt"
	"os"
)

var (
	// ErrNotSet is an error that is returned when the environment variable is not set.
	ErrNotSet = errors.New("not set")
	// ErrInvalidValue is an error that is returned when the environment variable is not valid.
	ErrInvalidValue = errors.New("invalid value")
)

// Env retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value will be parsed and returned.
// Otherwise, an error will be returned.
func Env[T Parsable](key string, options ...Option) (T, error) {
	var zero T

	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		return zero, fmt.Errorf("environment variable %q: %w", key, ErrNotSet)
	}

	val, err := parseAs[T](raw, newParams(options))
	if err != nil {
		return zero, fmt.Errorf("environment variable %q: %w", key, err)
	}

	return val, nil
}

// EnvOrDefault retrieves the value of the environment variable named by the key.
// If the variable is set and its value can be parsed, the parsed value is returned.
// Otherwise, defaultVal is returned: both when the variable is not set (or empty)
// and when its value is invalid. Use Env to tell these cases apart.
func EnvOrDefault[T Parsable](key string, defaultVal T, options ...Option) T {
	val, err := Env[T](key, options...)
	if err != nil {
		return defaultVal
	}

	return val
}
