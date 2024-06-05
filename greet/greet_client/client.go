package main

import (
	"fmt"
	"log"

	"github.com/godev99/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello gRPC client")
	cc, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC client cannot connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
