package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/kenriortega/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName

	reponse := greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return &reponse, nil
}

func (*server) GreetManyTimes(
	req *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer,
) error {
	for i := 0; i < 10; i++ {
		result := "Hello [" + strconv.Itoa(i) + "]"
		res := &greetpb.GreeManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error %v", err)
		}

		firstName := req.Greeting.FirstName
		result += "hi " + firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("err %v", err)
			return err
		}

		result := "hello " + req.Greeting.GetFirstName() + "!"

		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("err %v", err)
			return err
		}

	}
}

func main() {
	fmt.Println("Hello grpc server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds, sslErr := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.pem")
	if sslErr != nil {
		log.Fatalf("Failed to parse credentials: %v", sslErr)
		return
	}
	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
