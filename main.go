package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//Import the generated code
	pb "github.com/dn462/go-grpc-example/protos"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Ping service to the server
	pb.RegisterDataProcessServer(s, &pb.DataProcessServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Serve gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
