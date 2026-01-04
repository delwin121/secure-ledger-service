package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

// ChaosMiddleware adds artificial latency to requests based on SIMULATE_DELAY_MS env var.
func ChaosMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delayStr := os.Getenv("SIMULATE_DELAY_MS")
		if delayStr != "" {
			if delayMs, err := strconv.Atoi(delayStr); err == nil && delayMs > 0 {
				time.Sleep(time.Duration(delayMs) * time.Millisecond)
			}
		}
		next.ServeHTTP(w, r)
	})
}
