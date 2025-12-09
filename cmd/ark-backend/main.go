package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version   = "dev"
	commitSHA = "unknown"
	buildDate = "unknown"
)

const (
	defaultPort = "8080"
	defaultHost = "0.0.0.0"
)

func main() {
	log.Printf("Starting Ark Backend v%s (commit: %s, built: %s)", version, commitSHA, buildDate)

	// Create server
	addr := fmt.Sprintf("%s:%s", defaultHost, defaultPort)
	srv := &http.Server{
		Addr:         addr,
		Handler:      setupRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Backend listening on http://%s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down backend...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Backend shutdown error: %v", err)
		os.Exit(1)
	}

	log.Println("Backend stopped")
}

func setupRouter() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", handleHealth)

	// API version endpoint
	mux.HandleFunc("/api/version", handleVersion)

	// TODO: Add authentication middleware
	// TODO: Add CORS middleware
	// TODO: Add rate limiting
	// TODO: Add training endpoints
	// TODO: Add policy endpoints
	// TODO: Add audit endpoints

	return mux
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy","version":"%s"}`, version)
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"version":"%s","commit":"%s","buildDate":"%s"}`,
		version, commitSHA, buildDate)
}
