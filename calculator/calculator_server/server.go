package main

import (
	"context"
	"fmt"
	"github.com/gsafcik/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Calculator RPC: %v\n", req)

	sum := req.GetNumOne() + req.GetNumTwo()

	res := &calculatorpb.SumResponse{
		Result: sum,
	}

	return res, nil
}

func (*server) PrimeDecomposition(req *calculatorpb.PrimeDecompositionRequest, stream calculatorpb.CalculatorService_PrimeDecompositionServer) error {
	fmt.Printf("PrimeDecomposition function was invoked with %v\n", req)

	number := req.GetNum()
	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&calculatorpb.PrimeDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number/divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
