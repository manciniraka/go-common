package logger

import (
	"log/slog"
	"os"
)

func NewJSON() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)
}
