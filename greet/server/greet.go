package main

import (
	"context"
	"log"
	pb "github.com/katarzynakawala/grpc-go/greet/proto"
)

func ( s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was involved with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}