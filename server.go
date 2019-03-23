package main

import (
	"fmt"
	"github.com/jpreese/catfact/catfact"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const port = 1337
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on specified port: %v", err)
	}

	s := catfact.Server{}
	grpcServer := grpc.NewServer()

	catfact.RegisterCatFactServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start service: %s", err)
	} else {
		log.Printf("Service started.")
	}
}
