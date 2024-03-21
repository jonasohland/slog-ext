package slogext_test

import (
	"log/slog"
	"os"
	"testing"

	slogext "github.com/jonasohland/slog-ext/pkg/slog-ext"
)

func Test(t *testing.T) {

	logger := slog.New(slogext.NewHandler(os.Stderr, slogext.LevelTrace))

	slogext.Trace(logger, "Hi!")

}
