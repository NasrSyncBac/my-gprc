package main

import (
	"context"
	"fmt"
	"log"

	cal "github.com/Syncbak-Git/my-gprc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()
	c := cal.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created the client %f", c)
	doUnary(c)
}

func doUnary(c cal.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &cal.CalculateRequest{
		Calculator: &cal.Calculator{
			Num1: 17,
			Num2: 7,
		},
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculate RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res.Result)
}
