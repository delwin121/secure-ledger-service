package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func (w *ResponseWriterWrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		wrapper := &ResponseWriterWrapper{ResponseWriter: w, StatusCode: http.StatusOK}
		
		next.ServeHTTP(wrapper, r)
		
		duration := time.Since(start)
		
		slog.Info("request processed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapper.StatusCode,
			"duration", duration,
			"user_agent", r.UserAgent(),
		)
	})
}
