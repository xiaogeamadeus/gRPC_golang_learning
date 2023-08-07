package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

// Greet RPC Endpoints
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet funtion was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
