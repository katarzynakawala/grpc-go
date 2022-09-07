package main

import (
	"context"
	"io"
	"log"

	pb "github.com/katarzynakawala/grpc-go/calculator/proto"
)

func doCalculatePrimes(c pb.CalculatorServiceClient) {
	log.Println("doCalculatePrimes was invoked")


	req := &pb.CalcPrimesRequest{
		Digit: 120,
	}

	stream, err := c.CalculatePrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling CalculatePrimes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Calculating primes: %d\n", res.Result)
	}
}