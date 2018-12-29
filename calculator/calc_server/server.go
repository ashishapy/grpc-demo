package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ashishapy/grpc-go-course/calculator/calculatorpb"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Printf("Invoking Calculator service with request: %v \n", req)

	num1 := req.GetCalculator().GetNum1()
	num2 := req.GetCalculator().GetNum2()

	result := num1 + num2

	res := &calculatorpb.CalcResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello! I am Calculator_Server.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server error while listenning: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
