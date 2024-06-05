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

	req := &greetpb.GreetRequest{FirstName: "gab", LastName: "ivanes"}

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	request := greetpb.Greeting{"FirstName": "Bob", "LastName": "Bibi"}

	fmt.Println("call greet:", c.Greet)

}
