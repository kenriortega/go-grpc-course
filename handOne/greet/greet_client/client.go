package main

import (
	"context"
	"fmt"
	"log"

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
