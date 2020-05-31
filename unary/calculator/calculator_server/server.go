package main

import (
	"context"
	"fmt"
	"github.com/kentoje/grpc-test/unary/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n: ", req)
	result := req.GetFirstNumber() + req.GetSecondNumber()
	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed to listen: %v, err")
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}