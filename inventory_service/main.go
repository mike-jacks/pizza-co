package main

import (
	"log"
	"os"

	"github.com/mike_jacks/pizza_co/inventory_service/adapters"
	"github.com/mike_jacks/pizza_co/inventory_service/adapters/persistence"
	"github.com/mike_jacks/pizza_co/inventory_service/config"
	"github.com/mike_jacks/pizza_co/inventory_service/grpc"
)

func main() {

	db := config.InitDB()

	repo := persistence.NewInventoryRepository(db)

	inventoryServer := adapters.NewInventoryServer(repo)

	server := grpc.NewServer(inventoryServer)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
		os.Stdout.Sync()
	}

}
