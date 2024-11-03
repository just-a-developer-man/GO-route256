package logging

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	InitLogger(JSONHandler, slog.LevelDebug, io.Discard)
	assert.IsType(t, &customLoggerMiddleware{}, slog.Default().Handler())
	if got, ok := slog.Default().Handler().(*customLoggerMiddleware); ok {
		assert.IsType(t, &slog.JSONHandler{}, got.next)
	} else {
		t.Error("Invalid handler type")
	}
	assert.True(t, slog.Default().Enabled(context.Background(), slog.LevelDebug))

	InitLogger(TextHandler, slog.LevelInfo, io.Discard)
	assert.IsType(t, &customLoggerMiddleware{}, slog.Default().Handler())
	if got, ok := slog.Default().Handler().(*customLoggerMiddleware); ok {
		assert.IsType(t, &slog.TextHandler{}, got.next)
	} else {
		t.Error("Invalid handler type")
	}
	assert.False(t, slog.Default().Enabled(context.Background(), slog.LevelDebug))
}
