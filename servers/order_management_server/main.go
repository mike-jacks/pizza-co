package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/mike_jacks/pizza_co/clients/inventory"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/generated/order_management/v1"
	"github.com/mike_jacks/pizza_co/servers/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if os.Getenv("USE_ENV_FILE") != "false" {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Printf("No .env file found or error loading it: %v", err)
		} else {
			log.Println(".env file loaded successfully")
		}
	} else {
		log.Println("Skipping loading .env file as USE_ENV_FILEs set to false")
	}

	inventoryClientHost := os.Getenv("INVENTORY_SERVICE_HOST")
	// Create inventory client
	inventoryClient, err := inventory.CreateInventoryClient(inventoryClientHost, config.InventoryServerPort)
	if err != nil {
		log.Fatalf("Failed to create inventory client ; %v", err)
		os.Stdout.Sync()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.OrderManagementServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Stdout.Sync()
	}

	grpcServer := grpc.NewServer()
	orderManagementServer := NewOrderManagementServer(inventoryClient)
	order_management_v1_pb.RegisterOrderManagementServiceServer(grpcServer, orderManagementServer)
	reflection.Register(grpcServer)

	log.Printf("Server is running on port :%d...", config.OrderManagementServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		os.Stdout.Sync()
	}

}
