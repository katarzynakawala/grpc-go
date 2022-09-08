package main

import (
	"context"
	"log"

	pb "github.com/katarzynakawala/grpc-go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient){
	log.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream %v\n", err)
	}

	numbers := []int32{2, 4, 2, 2}

	for _, number := range numbers {
		log.Printf("Sending number %v\n", number)

		stream.Send(&pb.AverageRequest{
			Digit: number,
		})
	}

		res, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error while receiving response from doAverage: %v\n", err)
		}

		log.Printf("Average: %f\n", res.Result)
	}
