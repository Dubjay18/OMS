package main

import (
	common "github.com/Dubjay18/OMS-common.git"
	pb "github.com/Dubjay18/OMS-common.git/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	httpPort         = common.EnvString("HTTP_PORT", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer conn.Close()
	log.Println("Dialing orders service at ", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	h := NewHandler(c)
	h.registerRoutes(mux)

	log.Printf("Starting server on %s", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
