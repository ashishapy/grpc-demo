package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ashishapy/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello! I am Calculator GRPC Client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dialing: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Unary function call...")

	req := &calculatorpb.CalcRequest{
		Calculator: &calculatorpb.CalcInput{
			Num1: 3,
			Num2: 10,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while using Calculator service: %v", err)
	}

	fmt.Printf("Sum of two numbers: %v \n", res)
}
