package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/greet/greetpb"
	"log"
)

func main() {
	fmt.Println("I am a client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f" , c)
}