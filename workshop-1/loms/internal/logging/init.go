package logging

import (
	"io"
	"log/slog"
)

type handlerType uint8

const (
	JSONHandler handlerType = iota
	TextHandler handlerType = 1
)

// InitLogger function initializes default slog.Logger with given params 
func InitLogger(hType handlerType, level slog.Level, writers ...io.Writer) {
	slog.Default().Handler()
	var defaultHandler slog.Handler
	opts := &slog.HandlerOptions{Level: level}
	multiWriter := io.MultiWriter(writers...)

	switch hType {
	case JSONHandler:
		defaultHandler = slog.NewJSONHandler(multiWriter, opts)
	case TextHandler:
		defaultHandler = slog.NewTextHandler(multiWriter, opts)
	}

	slog.SetDefault(slog.New(newCustomLogger(defaultHandler)))
}
