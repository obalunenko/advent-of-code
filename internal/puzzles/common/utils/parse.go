package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseInts parses io.Reader into []int.
func ParseInts(in io.Reader, sep string) ([]int, error) {
	scanner := bufio.NewScanner(in)

	var res []int

	for scanner.Scan() {
		line := scanner.Text()

		if sep == "" {
			n, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("parse int: %w", err)
			}

			res = append(res, n)
		} else {
			split := strings.Split(line, sep)

			for _, s := range split {
				n, err := strconv.Atoi(s)
				if err != nil {
					return nil, fmt.Errorf("parse int: %w", err)
				}

				res = append(res, n)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return res, nil
}
