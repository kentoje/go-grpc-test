package main

import (
	"context"
	"fmt"
	"github.com/kentoje/grpc-test/streaming-server/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("I am a client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewPrimeNumberServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.PrimeNumberServiceClient)  {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 210,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberManyTimes RPC: %v", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		fmt.Printf("Response from GreetManyTimes: %v\n", res.GetPrimeFactor())
	}
}
