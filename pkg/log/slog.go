package log

import (
	"log/slog"
	"os"
	"strings"
)

type slogLogger struct {
	h *slog.Logger
}

func (slogger *slogLogger) InitLogger(levelStr string) {
	logConfig := &slog.HandlerOptions{
		AddSource: false,
	}
	level := slogger.getLogLevelFromString(levelStr)
	logConfig.Level = level
	slogger.h = slog.New(slog.NewJSONHandler(os.Stdout, logConfig))
}

func (slogger *slogLogger) Debug(msg string, args ...any) {
	slogger.h.Debug(msg, args...)
}

func (slogger *slogLogger) Info(msg string, args ...any) {
	slogger.h.Info(msg, args...)
}

func (slogger *slogLogger) Warn(msg string, args ...any) {
	slogger.h.Warn(msg, args...)
}

func (slogger *slogLogger) Error(msg string, args ...any) {
	slogger.h.Error(msg, args...)
}

func (slogger *slogLogger) getLogLevelFromString(levelStr string) slog.Level {
	levelStr = strings.ToLower(strings.TrimSpace(levelStr))

	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo // Default to info
	}
}
