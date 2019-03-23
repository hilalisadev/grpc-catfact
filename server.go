package main

import (
	"fmt"
	"github.com/jpreese/catfact/catfact"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const port = 7777

	log.Printf("Attempting to listen on port: %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d (%v)", port, err)
	}
	defer lis.Close()

	catFactServer := catfact.Server{}
	grpcServer := grpc.NewServer()
	log.Printf("Service created.")

	catfact.RegisterCatFactServiceServer(grpcServer, &catFactServer)
	log.Printf("Service registered.")

	log.Printf("Attempting to start service.")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start service: %s", err)
	} else {
		log.Printf("Service started.")
	}
}
