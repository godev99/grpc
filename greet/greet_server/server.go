package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/godev99/grpc/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", listener)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	fmt.Println("Hello gRPC server !")
}
