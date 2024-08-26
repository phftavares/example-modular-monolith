package logger

import (
	"context"

	"go.uber.org/zap"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func getLogger(ctx context.Context) *zap.Logger {
	logger, _ := ctx.Value(loggerKey{}).(*zap.Logger)
	return logger
}
