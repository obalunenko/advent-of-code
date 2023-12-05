// Package templates contains templates for solution.go, solution_test.go and spec.md files.
package templates

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
)

var (
	//go:embed solution.go.tmpl
	solutionTmpl string
	//go:embed solution_test.go.tmpl
	solutionTestTmpl string
	//go:embed spec.md.tmpl
	specTmpl string
)

// Params contains parameters for templates.
type Params struct {
	Year   string // e.g. "2023"
	Day    int    // e.g. 2
	DayStr string // e.g. "02"
	URL    string // e.g. "https://adventofcode.com/2023/day/2"
}

// SolutionTmpl returns template for solution.go file.
func SolutionTmpl() (*template.Template, error) {
	tmpl, err := template.New("solution.go").Parse(solutionTmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solution template: %w", err)
	}

	return tmpl, nil
}

// SolutionTestTmpl returns template for solution_test.go file.
func SolutionTestTmpl() (*template.Template, error) {
	tmpl, err := template.New("solution_test.go").Parse(solutionTestTmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse solution test template: %w", err)
	}

	return tmpl, nil
}

// SpecTmpl returns template for spec.md file.
func SpecTmpl() (*template.Template, error) {
	tmpl, err := template.New("spec.md").Parse(specTmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec template: %w", err)
	}

	return tmpl, nil
}

// SubstituteTemplate substitutes template with given parameters.
func SubstituteTemplate(tmpl *template.Template, p Params) ([]byte, error) {
	var buf bytes.Buffer

	err := tmpl.Execute(&buf, p)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}
