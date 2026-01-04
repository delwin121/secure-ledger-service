package main

import (
	"ledger-service/handlers"
	"ledger-service/middleware"
	"ledger-service/store"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Initialize dependencies
	store := store.NewTransactionStore()
	handler := handlers.NewTransactionHandler(store)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/transaction", handler.Create)
	mux.Handle("/metrics", promhttp.Handler())

	// Add middleware (wrapping one after another)
	// Order: Logging -> Metrics -> Chaos -> Mux
	var srvHandler http.Handler = mux
	srvHandler = middleware.ChaosMiddleware(srvHandler)
	srvHandler = middleware.MetricsMiddleware(srvHandler)
	srvHandler = middleware.LoggingMiddleware(srvHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Ensure port starts with :
	if len(port) > 0 && port[0] != ':' {
		port = ":" + port
	}

	slog.Info("server starting", "port", port)
	if err := http.ListenAndServe(port, srvHandler); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
