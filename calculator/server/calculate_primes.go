package main

import (
	"log"

	pb "github.com/katarzynakawala/grpc-go/calculator/proto"
)

func (s *Server) CalculatePrimes(in *pb.CalcPrimesRequest, stream pb.CalculatorService_CalculatePrimesServer) error {
	log.Printf("CalculatePrimes function was invoked with %v\n", in)

	//var results []int32
	var k int32 = 2
	var number int32 = in.Digit

	for number > 1 {
		if number % k == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			number /= k 
		} else {
			k++
		}
		
	}
	return nil
	} 

