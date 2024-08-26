package logger

import (
	"context"

	"go.uber.org/zap"
)

func Debug(ctx context.Context, message string, fields ...zap.Field) {
	l := getLogger(ctx)
	if l == nil {
		return
	}

	l.Debug(message, fields...)
}

func Info(ctx context.Context, message string, fields ...zap.Field) {
	l := getLogger(ctx)
	if l == nil {
		return
	}

	l.Info(message, fields...)
}

func Error(ctx context.Context, message string, fields ...zap.Field) {
	l := getLogger(ctx)
	if l == nil {
		return
	}

	l.Error(message, fields...)
}

func Warn(ctx context.Context, message string, fields ...zap.Field) {
	l := getLogger(ctx)
	if l == nil {
		return
	}

	l.Warn(message, fields...)
}
