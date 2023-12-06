package main

import (
	"log"
	"net"
	"github.com/abdullahchand/gRPC-Protobuf/golang/proto"
	"github.com/abdullahchand/gRPC-Protobuf/golang/commands"
	"google.golang.org/grpc"
)



func main() {
	lis, err :=net.Listen("tcp", ":5002");
	if err != nil{
		log.Fatalf("Couldnt create a TCP socket on port 5002: %s", err)
	}

	// Setup gRPC server.
	grpc_server := grpc.NewServer()
	
	// Point to your type that implements the functions from proto.
	service := &commands.MathServer{}

	proto.RegisterAdditionServer(grpc_server, service)

	err = grpc_server.Serve(lis)

	if err != nil {
		log.Fatalf("Couldnt server : %s", err)
	}
}