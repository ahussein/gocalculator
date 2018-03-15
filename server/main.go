// main server file

package main

import (
	"log"
	"net"

	"github.com/ahussein/gocalculator/pb"
	"github.com/ahussein/gocalculator/calculator"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const port = ":9000"

// main works as the entry point for starting the calculator service
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


// CalculateorService implements the Calculator interface
type CalculatorService struct{}

// Add adds tow integers and return a float result
func (c *CalculatorService) Add(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	result := calculator.Add(params.Param1, params.Param2)
	return &pb.Result{Result: result}, nil
}

// Subtract substracts two integers and return a float result
func (c *CalculatorService) Subtract(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	result := calculator.Subtract(params.Param1, params.Param2)
	return &pb.Result{Result: result}, nil
}

// Multiply multiplies two integers and return a float result
func (c *CalculatorService) Multiply(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	result := calculator.Multiply(params.Param1, params.Param2)
	return &pb.Result{Result: result}, nil
}

// Divide divides two integers and return a float result
func (c *CalculatorService) Divide(ctx context.Context, params *pb.Parameters) (*pb.Result, error) {
	result, err := calculator.Divide(params.Param1, params.Param2)
	return &pb.Result{Result: result}, err
}
