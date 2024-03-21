package slogext

import (
	"context"
	"log/slog"
)

type LoggerContextKeyType int

const LoggerContextKey LoggerContextKeyType = iota

func WithContext(ctx context.Context, log *slog.Logger) context.Context {
	return context.WithValue(ctx, LoggerContextKey, log)
}

func FromContext(ctx context.Context) *slog.Logger {
	v := ctx.Value(LoggerContextKey)
	if v == nil {
		return slog.Default()
	}

	logger, ok := v.(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return logger
}
