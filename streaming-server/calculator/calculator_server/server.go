package main

import (
	"fmt"
	"github.com/kentoje/grpc-test/streaming-server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.PrimeNumberService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberManyTimes function was invoked with %v\n", req)

	k := int64(2)
	number := req.GetNumber()

	for number > 1 {
		if number % k == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: k,
			})

			number = number / k
		} else {
			k++
			fmt.Printf("Divisor has increased to: %v\n", k)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterPrimeNumberServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
