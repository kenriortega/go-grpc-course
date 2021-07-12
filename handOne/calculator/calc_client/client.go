package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// doUnary(c)
	// doStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
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
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("err %v", err)
	}

	numbers := []int32{1, 23, 4, 5, 6, 677, 2}

	for _, number := range numbers {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Println("Result ", res.GetAverage())
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {

	stream, err := c.FindMaximun(context.Background())
	if err != nil {
		log.Fatalf("err %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		numbers := []int32{3, 4, 55, 67, 8, 23}
		for _, v := range numbers {
			stream.Send(&calculatorpb.FindMaximunRequest{
				Number: v,
			})
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
				log.Fatalf("err: %v", err)
				break
			}
			maximun := res.GetMaximun()
			fmt.Println(maximun)
		}
		close(waitc)
	}()
	<-waitc
}
