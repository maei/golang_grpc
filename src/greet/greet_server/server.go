package main

import (
	"fmt"
	"github.com/maei/golang_grpc/src/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

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
