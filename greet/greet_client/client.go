package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/kenriortega/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello from client")
	creds, sslErr := credentials.NewClientTLSFromFile("./ssl/ca.crt", "")
	if sslErr != nil {
		log.Fatalf("Failed to parse credentials: %v", sslErr)
		return
	}
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	// doUnaryWhithDeadline(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)
}

func doUnaryWhithDeadline(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ka",
			LastName:  "as",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalf("Error: %v", statusErr.Message())
			} else {
				log.Fatalf("Error: %v", statusErr)
			}
		} else {
			log.Fatalf("Error: %v", err)
		}
	}
	fmt.Println(res)
}

func doUnary(c greetpb.GreetServiceClient) {
	res, err := c.Greet(context.Background(),
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "kalix",
				LastName:  "Ortega",
			},
		},
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "kali",
			LastName:  "ort",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err %v", err)
		}
		log.Printf("Response for greetmanytimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ka",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ke",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ki",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ko",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ku",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("err %v", err)
	}

	for _, req := range requests {

		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err %v", err)
	}

	fmt.Printf("response %v", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ka",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ke",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ki",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ko",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "ku",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("err %v", err)
		return
	}
	waitc := make(chan struct{})
	go func() {
		for _, v := range requests {
			stream.Send(v)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("err %v", err)
				break
			}
			fmt.Println(res.GetResult())
		}
		close(waitc)

	}()
	<-waitc
}
