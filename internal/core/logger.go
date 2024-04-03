package core

import (
	"io"
	"log/slog"
)

func NewLogger(out io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(out, nil))
}
