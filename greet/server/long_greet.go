package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/katarzynakawala/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked.")

	var res string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving the request %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)

	}
}