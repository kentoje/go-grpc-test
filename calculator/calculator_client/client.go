package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/calculator/calculatorpb"
	"log"
)

func main() {
	fmt.Println("Here is the client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting to do Unary RPC...")

	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Calculator: %v", res.Result)
}