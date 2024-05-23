package main

import (
	"context"
	common "github.com/Dubjay18/OMS-common.git"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()
	store := NewStore()
	svc := NewService(store)
	NewGrpcHandler(grpcServer, svc)
	svc.CreateOrder(context.Background())

	log.Println("Starting server on ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
