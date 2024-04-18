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

			if serverType == "r" && r.Method != http.MethodGet {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Mirror is read-only"))
				w.WriteHeader(http.StatusForbidden)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
