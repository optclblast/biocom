package logger

import (
	"log/slog"
	"os"
)

// NewLocalLogger returns a pretty slog logger for local development.
func NewLocalLogger() *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

// NewLogger returns slog Logger with JSONHandler
func NewLogger(lvl slog.Level) *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}),
	)
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func MapLevel(env string) slog.Level {
	switch env {
	case "dev", "local", "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
