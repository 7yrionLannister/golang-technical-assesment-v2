package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	h *zap.Logger
}

func (zlogger *zapLogger) InitLogger(levelStr string) {
	levelStr = strings.ToLower(strings.TrimSpace(levelStr))

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	if levelStr != "info" {
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		config.Development = true
	}

	l := zap.Must(config.Build())
	zlogger.h = l.WithOptions(zap.AddCallerSkip(1)) // Skip the logger.go file
}

func (zlogger *zapLogger) Debug(msg string, args ...any) {
	zapFields := zlogger.anySliceToZapFieldSlice(args)
	zlogger.h.Debug(msg, zapFields...)
}

func (zlogger *zapLogger) Info(msg string, args ...any) {
	zapFields := zlogger.anySliceToZapFieldSlice(args)
	zlogger.h.Info(msg, zapFields...)
}

func (zlogger *zapLogger) Warn(msg string, args ...any) {
	zapFields := zlogger.anySliceToZapFieldSlice(args)
	zlogger.h.Warn(msg, zapFields...)
}

func (zlogger *zapLogger) Error(msg string, args ...any) {
	zapFields := zlogger.anySliceToZapFieldSlice(args)
	zlogger.h.Error(msg, zapFields...)
}

func (zlogger *zapLogger) anySliceToZapFieldSlice(fields []any) []zap.Field {
	var zapFields []zap.Field
	n := len(fields)
	for i := 0; i < n; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue // Skip if key is not a string
		}

		switch v := fields[i+1].(type) {
		case string:
			zapFields = append(zapFields, zap.String(key, v))
		case int:
			zapFields = append(zapFields, zap.Int(key, v))
		case int64:
			zapFields = append(zapFields, zap.Int64(key, v))
		case float64:
			zapFields = append(zapFields, zap.Float64(key, v))
		case bool:
			zapFields = append(zapFields, zap.Bool(key, v))
		default:
			zapFields = append(zapFields, zap.Any(key, v))
		}
	}
	return zapFields
}
