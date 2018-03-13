package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ahussein/gocalculator/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":9000"

func main() {
	param1 := flag.Int("Param1", 0, "First parameter")
	param2 := flag.Int("Param2", 0, "Second parameter")
	address := flag.String("Address", "localhost:9000", "Address to the calculator server")
	operation := flag.String("Operation", "add", "Operation to execute, one of (add, subtract, multiply, divide)")
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// close connection later
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	switch *operation {
	case "add":
		// need to find away to avoid duplicate the error handling and printing the result
		res, err := client.Add(context.Background(),
			&pb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "subtract":
		res, err := client.Subtract(context.Background(),
			&pb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "multiply":
		res, err := client.Multiply(context.Background(),
			&pb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "divide":
		res, err := client.Divide(context.Background(),
			&pb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	default:
		fmt.Println("Operation must be one of the following (add, subtract, multiply, divide)")
		os.Exit(1)
	}
}
