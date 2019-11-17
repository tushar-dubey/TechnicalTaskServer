package main

import (
	"TechnicalTaskServer/calculator/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (c *server) SumIntegers(ctx context.Context, sumRequest *proto.SumRequest) (*proto.SumResponse, error) {
	fmt.Printf("sum service was invoked : %v", sumRequest)
	sum := sumRequest.Int1 + sumRequest.Int2
	result := &proto.SumResponse{
		Result:               sum,
	}

	return result, nil
}

func (c *server) SubtractIntegers(ctx context.Context, subtractRequest *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	fmt.Printf("subtract service was invoked : %v", subtractRequest)
	subtract := subtractRequest.Int1 - subtractRequest.Int2
	result := &proto.SubtractResponse{
		Result:               subtract,
	}

	return result, nil
}

func main() {
	fmt.Println("Welcome to Simple Calculator Service")

	lis, err := net.Listen("tcp", "127.0.0.1:50051")

	if err!= nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve , %v", err)
	}
}
