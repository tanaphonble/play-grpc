package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %+v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return res, nil
}

func (s *server) PrimeNumberDecomosition(req *calculatorpb.PrimeNumberDecomositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecomositionServer) error {
	fmt.Printf("Received PrimeNumberDecomosition RPC: %+v", req)
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecomositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has  incresed to %d", divisor)
		}
	}

	return nil
}

func main() {
	protocol, address := "tcp", "0.0.0.0:50051"
	lis, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatalf("Failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
