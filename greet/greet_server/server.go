package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/rwoj/my-grpc/greet/greetpb"
)


func main() {
	type myserver struct{}
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &myserver{})

	if err:= s.Serve(lis); err !=nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
