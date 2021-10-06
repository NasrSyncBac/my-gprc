package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Syncbak-Git/my-gprc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %s", cc)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nasr",
			LastName:  "Mohammed",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %s", err)
	}

	log.Printf("Response from Greet: %s", res.Result)
}
