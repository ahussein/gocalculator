package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ahussein/gocalculator/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":9000"

func main() {
	param1 := flag.Int("Param1", 0, "First parameter")
	param2 := flag.Int("Param2", 0, "Second parameter")
	address := flag.String("Address", "localhost:9000", "Address to the calculator server")
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// close connection later
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Add(context.Background(),
		&pb.Parameters{Param1: int32(*param1),
			Param2: int32(*param2)})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Result)
}
