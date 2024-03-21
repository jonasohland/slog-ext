package slogext

import (
	"io"
	"log/slog"

	console "github.com/phsym/console-slog"
)

func NewHandler(io io.Writer, level slog.Leveler) slog.Handler {
	return console.NewHandler(io, &console.HandlerOptions{Level: level})
}
