package middleware

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

func NewAccessLog(serverType string, log *zap.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.Background())

			log.Info("New request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("protocol", r.Proto),
				zap.String("origin", r.Header.Get("Origin")))

			handler.ServeHTTP(w, r)
		})
	}
}
