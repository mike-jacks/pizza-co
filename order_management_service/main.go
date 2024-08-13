package main

import (
	"log"
	"os"

	"github.com/mike_jacks/pizza_co/order_management_service/grpc"
)

func main() {
	if err := grpc.GetENV(); err != nil {
		log.Println(err)
		os.Stdout.Sync()
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
		os.Stdout.Sync()
	}
}
