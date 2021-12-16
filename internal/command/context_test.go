package command_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/advent-of-code/internal/command"
	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

func TestContext(t *testing.T) {
	sess := "test_session"

	ctx := command.ContextWithSession(context.Background(), sess)

	got := command.SessionFromContext(ctx)
	assert.Equal(t, sess, got)

	got = command.SessionFromContext(nil)
	assert.Equal(t, "", got)

	got = command.SessionFromContext(context.Background())
	assert.Equal(t, "", got)

	opt := puzzles.WithElapsed()

	ctx = command.ContextWithOptions(ctx, opt)

	gotopts := command.OptionsFromContext(ctx)
	assert.Equal(t, []puzzles.RunOption{opt}, gotopts)

	gotopts = command.OptionsFromContext(nil)
	assert.Equal(t, []puzzles.RunOption{}, gotopts)

	gotopts = command.OptionsFromContext(context.Background())
	assert.Equal(t, []puzzles.RunOption{}, gotopts)
}
