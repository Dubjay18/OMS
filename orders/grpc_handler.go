package main

import (
	"context"
	pb "github.com/Dubjay18/OMS-common.git/api"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return &grpcHandler{}
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order recieved!")
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
