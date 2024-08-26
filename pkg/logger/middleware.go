package logger

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func newLogger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("GO_ENVIRONMENT") == "production" {
		logger = zap.Must(zap.NewProduction())
	}
	return logger
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		l := newLogger()
		l.With(String("request_id", middleware.GetReqID(ctx)))

		ctx = WithLogger(ctx, l)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
