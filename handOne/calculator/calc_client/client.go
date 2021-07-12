package main

import (
	"context"
	"fmt"
	"io"
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
	doStreaming(c)
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

func doStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("Recv msg %v\n", res.GetPrimerFactor())
	}
}
