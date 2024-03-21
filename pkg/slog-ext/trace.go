package slogext

import (
	"context"
	"log/slog"
)

const LevelTrace slog.Level = -8

type traceLeveler struct {
}

func (t *traceLeveler) Level() slog.Level {
	return LevelTrace
}

func Trace(log *slog.Logger, msg string, args ...any) {
	log.Log(context.Background(), LevelTrace, msg, args...)
}

func TraceLeveler() slog.Leveler {
	return &traceLeveler{}
}
