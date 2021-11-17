package puzzles

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

// Result represents puzzle solution result.
type Result struct {
	Year    string
	Name    string
	Part1   string
	Part2   string
	metrics metrics
}

func (r Result) String() string {
	if r.Part1 == "" {
		r.Part1 = unsolved
	}

	if r.Part2 == "" {
		r.Part2 = unsolved
	}

	if r.Name == "" {
		r.Name = unknown
	}

	if r.Year == "" {
		r.Year = unknown
	}

	var buf strings.Builder

	const linesnum = 10

	table := make([][]string, 0, linesnum)

	table = append(table, []string{})
	puzzleHeaderLine := []string{fmt.Sprintf("%s/%s puzzle answer:", r.Year, r.Name)}
	table = append(table, puzzleHeaderLine)

	part1Line := []string{"part1", r.Part1}
	part2Line := []string{"part2", r.Part2}

	table = append(table, part1Line, part2Line)
	table = append(table, []string{})

	if r.metrics != nil {
		metricsHeaderLine := []string{"metrics:"}

		table = append(table, metricsHeaderLine)

		for i := range r.metrics {
			m := r.metrics[i]

			var line []string

			if m.mType.HasFlag(metricsFlagNone) {
				line = []string{m.mType.String()}
			} else {
				line = []string{m.mType.String(), m.metadata}
			}

			table = append(table, line)
		}
	}
	table = append(table, []string{})
	if err := printTable(&buf, table); err != nil {
		panic(err)
	}

	content := strings.TrimSpace(buf.String())

	content = "\n" + "   " + content + "\n"

	return content
}

func printTable(w io.Writer, table [][]string) error {
	writer := tabwriter.NewWriter(w, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)

	var (
		err error
	)

	for _, line := range table {
		switch len(line) {
		case 0:
			_, err = fmt.Fprintf(w, "\n")
		case 1:
			_, err = fmt.Fprintf(writer, "\t"+strings.Join(line, "\t")+"\t\n")
		default:
			_, err = fmt.Fprintf(writer, "\t   "+strings.Join(line, "\t")+"\t\n")
		}

		if err != nil {
			if err != nil {
				return fmt.Errorf("fprintln: %w", err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush: %w", err)
	}

	return nil
}
