package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kenriortega/go-grpc-course/handOne/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client calculator")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling sum rpc: %v", err)
	}
	log.Println("response: ", res)
}
