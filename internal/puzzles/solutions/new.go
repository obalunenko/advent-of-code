package solutions

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/obalunenko/advent-of-code/internal/puzzles/solutions/templates"
)

func createNewFromTemplate(purl string) error {
	const (
		perms   = 0o655
		yearLen = 4
		dayLen  = 2
	)

	pd, err := parsePuzzleURL(purl)
	if err != nil {
		return fmt.Errorf("parse puzzle url %q: %w", purl, err)
	}

	day := strconv.Itoa(pd.day)
	if len(day) < dayLen {
		day = "0" + day
	}

	if len(day) != dayLen {
		return fmt.Errorf("invalid day: %s", day)
	}

	year := strconv.Itoa(pd.year)

	if len(year) != yearLen {
		return fmt.Errorf("invalid year: %s", year)
	}

	params := templates.Params{
		Year:   year,
		Day:    pd.day,
		DayStr: day,
		URL:    purl,
	}

	path := filepath.Clean(filepath.Join(year, "day"+day))

	if err = createPuzzleDir(path, perms); err != nil {
		return fmt.Errorf("failed to create puzzle dir: %w", err)
	}

	testdata := filepath.Clean(filepath.Join(path, "testdata"))

	if err = createTestdata(testdata, perms); err != nil {
		return fmt.Errorf("failed to create testdata: %w", err)
	}

	tmplsFns := []func() (*template.Template, error){
		templates.SolutionTmpl, templates.SolutionTestTmpl, templates.SpecTmpl,
	}

	for _, tmplFn := range tmplsFns {
		var tmpl *template.Template

		tmpl, err = tmplFn()
		if err != nil {
			return fmt.Errorf("failed to get template: %w", err)
		}

		if err = createFromTemplate(tmpl, path, perms, params); err != nil {
			return fmt.Errorf("failed to create from template: %w", err)
		}
	}

	return nil
}

func createFromTemplate(tmpl *template.Template, path string, perms os.FileMode, params templates.Params) error {
	fpath := filepath.Clean(filepath.Join(path, tmpl.Name()))

	if isExist(fpath) {
		return nil
	}

	var content []byte

	content, err := templates.SubstituteTemplate(tmpl, params)
	if err != nil {
		return fmt.Errorf("failed to substitute template: %w", err)
	}

	if err = os.WriteFile(fpath, content, perms); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func createPuzzleDir(path string, perms os.FileMode) error {
	if !isExist(path) {
		if err := os.MkdirAll(path, perms); err != nil {
			return fmt.Errorf("failed to create dir: %w", err)
		}
	}

	return nil
}

func createTestdata(path string, perms os.FileMode) error {
	if !isExist(path) {
		if err := os.MkdirAll(path, perms); err != nil {
			return fmt.Errorf("failed to create dir: %w", err)
		}
	}

	input := filepath.Clean(filepath.Join(path, "input.txt"))

	if !isExist(input) {
		var f *os.File

		f, err := os.Create(input)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		if err = f.Close(); err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}
	}

	return nil
}

func isExist(path string) bool {
	stat, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return false
	}

	return stat != nil && stat.Name() != ""
}

type puzzleDate struct {
	year int
	day  int
}

func parsePuzzleURL(url string) (puzzleDate, error) {
	const (
		urlFmt    = "https://adventofcode.com/%d/day/%d"
		paramsNum = 2
	)

	var year, day int

	n, err := fmt.Sscanf(url, urlFmt, &year, &day)
	if err != nil {
		return puzzleDate{}, fmt.Errorf("parse puzzle url: %w", err)
	}

	if n != paramsNum {
		return puzzleDate{}, fmt.Errorf("invalid puzzle url: %s", url)
	}

	return puzzleDate{
		year: year,
		day:  day,
	}, nil
}
