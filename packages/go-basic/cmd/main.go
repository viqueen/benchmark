package main

import (
	"api/go-sdk/content/v1/contentV1connect"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"packages/go-basic/internal/api"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	contentService := api.NewContentService()
	contentPath, contentHandler := contentV1connect.NewContentServiceHandler(contentService)

	mux := http.NewServeMux()
	mux.Handle(contentPath, contentHandler)

	address := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine so it doesn't block
	go func() {
		log.Printf("Starting server on %s", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Block until we receive a signal
	<-stop
	log.Println("Shutting down server...")

	// Create a deadline for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
