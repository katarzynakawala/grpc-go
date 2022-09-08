package main

import (
	"io"
	"log"

	pb "github.com/katarzynakawala/grpc-go/calculator/proto"
)

func (s * Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")

	var res int32
	counter := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(res) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving number %v\n", req.Digit)


		res += req.Digit
		counter++
	}

}



	