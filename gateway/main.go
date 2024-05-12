package main

import (
	"log"
	"net/http"
)

const (
	httpPort = ":8080"
)

func main() {
	mux := http.NewServeMux()
	h := NewHandler()
	h.registerRoutes(mux)

	log.Printf("Starting server on %s", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
