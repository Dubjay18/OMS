package main

import (
	common "github.com/Dubjay18/OMS-common.git"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

var (
	httpPort = common.EnvString("HTTP_PORT", ":8080")
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
