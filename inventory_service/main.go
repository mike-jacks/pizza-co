package main

import (
	"log"
	"os"

	"github.com/mike_jacks/pizza_co/inventory_service/grpc"
)

func main() {

	server := grpc.NewServer()

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
		os.Stdout.Sync()
	}

}
