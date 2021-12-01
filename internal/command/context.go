package command

import (
	"context"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

type sessCtxKey struct{}

// ContextWithSession adds session to context.
func ContextWithSession(ctx context.Context, session string) context.Context {
	if session == "" {
		return ctx
	}

	return context.WithValue(ctx, sessCtxKey{}, session)
}

// SessionFromContext extracts session from context.
func SessionFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	val := ctx.Value(sessCtxKey{})
	sess, ok := val.(string)
	if !ok {
		return ""
	}

	return sess
}

type optsCtxKey struct{}

// ContextWithOptions adds options to context.
func ContextWithOptions(ctx context.Context, opts []puzzles.RunOption) context.Context {
	if len(opts) == 0 {
		return ctx
	}

	return context.WithValue(ctx, optsCtxKey{}, opts)
}

// OptionsFromContext extracts options from context.
func OptionsFromContext(ctx context.Context) []puzzles.RunOption {
	if ctx == nil {
		return []puzzles.RunOption{}
	}

	v := ctx.Value(optsCtxKey{})

	opts, ok := v.([]puzzles.RunOption)
	if !ok {
		return []puzzles.RunOption{}
	}

	return opts
}
