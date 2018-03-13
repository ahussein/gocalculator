// main server file

package main

import (
	"errors"
	"log"
	"net"

	"github.com/ahussein/gocalculator/pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	// for this simple task, we do not need to secure the communication
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServer(s, new(CalculatorService))
	log.Println("Starting server on port " + port)
	s.Serve(lis)

}

type CalculatorService struct{}

func (c *CalculatorService) Add(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	var result int32
	result = params.Param1 + params.Param2
	return &pb.Result{Result: float32(result)}, nil
}

func (c *CalculatorService) Subtract(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	var result int32
	result = params.Param1 - params.Param2
	return &pb.Result{Result: float32(result)}, nil
}

func (c *CalculatorService) Multiply(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	var result int32
	result = params.Param1 * params.Param2
	return &pb.Result{Result: float32(result)}, nil
}

func (c *CalculatorService) Divide(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	if params.Param2 == 0 {
		return nil, errors.New("Division by zero is not allowed")
	}
	var result int32
	result = params.Param1 / params.Param2
	return &pb.Result{Result: float32(result)}, nil
}
