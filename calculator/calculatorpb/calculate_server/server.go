package main

import (
	"context"
	"fmt"
	"log"
	"net"

	cal "github.com/Syncbak-Git/my-gprc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	cal.UnimplementedCalculatorServiceServer
}

func (*server) Calculate(ctx context.Context, req *cal.CalculateRequest) (*cal.CalculateResponse, error) {
	fmt.Printf("Calculate function was invoked with %v\n", req)
	num1 := req.GetCalculator().GetNum1()
	num2 := req.GetCalculator().GetNum2()
	result := num1 + num2
	res := &cal.CalculateResponse{
		Result: fmt.Sprint(result),
	}
	return res, nil
}
func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cal.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
