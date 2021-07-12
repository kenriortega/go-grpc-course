package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/kenriortega/go-grpc-course/handOne/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
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
