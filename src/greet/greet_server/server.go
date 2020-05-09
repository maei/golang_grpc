package main

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc/src/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	// We want the first_name and last_name
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	name := fmt.Sprintf(firstName + " " + lastName)

	result := fmt.Sprintf("Hello %s", name)
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello from the Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("Failed to list: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	errServ := s.Serve(lis)
	if errServ != nil {
		log.Fatalf("failed to serve: %v", errServ)
	}
}
