package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mike_jacks/pizza_co/clients/inventory"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/generated/order_management/v1"
	"github.com/mike_jacks/pizza_co/servers/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	inventoryClientHost := os.Getenv("INVENTORY_SERVICE_HOST")
	// Create inventory client
	inventoryClient, err := inventory.CreateInventoryClient(inventoryClientHost, config.InventoryServerPort)
	if err != nil {
		log.Fatalf("Failed to create inventory client ; %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.OrderManagementServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderManagementServer := NewOrderManagementServer(inventoryClient)
	order_management_v1_pb.RegisterOrderManagementServiceServer(grpcServer, orderManagementServer)
	reflection.Register(grpcServer)

	log.Printf("Server is running on port :%d...", config.OrderManagementServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
