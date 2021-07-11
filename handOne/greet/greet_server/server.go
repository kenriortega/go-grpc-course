package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kenriortega/go-grpc-course/handOne/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	reponse := greetpb.GreetResponse{
		Result: result,
	}

	return &reponse, nil
}

func main() {
	fmt.Println("Hello grpc server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
