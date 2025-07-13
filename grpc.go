package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *grpcServer {
	return &grpcServer{
		addr: addr,
	}
}

func (s *grpcServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("Starting the gRPC server on ", s.addr)

	return grpcServer.Serve(lis)
}
