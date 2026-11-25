package getenv

import (
	"time"
)

// Defaults for the options that were not passed.
const (
	defaultSeparator  = ","
	defaultTimeLayout = time.RFC3339
)

// Option configures how Env and EnvOrDefault parse the value of a variable.
type Option interface {
	apply(p *params)
}

// params holds the parameters that options set.
type params struct {
	// separator splits slice values.
	separator string
	// layout parses time.Time values.
	layout string
}

type withSeparator string

func (w withSeparator) apply(p *params) {
	p.separator = string(w)
}

// WithSeparator sets the separator of slice values. The default is ",".
func WithSeparator(separator string) Option {
	return withSeparator(separator)
}

type withTimeLayout string

func (w withTimeLayout) apply(p *params) {
	p.layout = string(w)
}

// WithTimeLayout sets the layout of time.Time values. The default is time.RFC3339.
func WithTimeLayout(layout string) Option {
	return withTimeLayout(layout)
}

// newParams creates new parameters from the defaults and options.
func newParams(opts []Option) params {
	p := params{
		separator: defaultSeparator,
		layout:    defaultTimeLayout,
	}

	for _, opt := range opts {
		opt.apply(&p)
	}

	return p
}
