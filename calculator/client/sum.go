package main

import (
	"context"
	"log"

	pb "github.com/katarzynakawala/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstDigit: 1,
		SecondDigit: 1,
	})

	if err != nil {
		log.Fatalf("Could not sum %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}