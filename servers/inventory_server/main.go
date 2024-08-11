package main

import (
	"fmt"
	"log"
	"net"

	inventory_v1_pb "github.com/mike_jacks/pizza_co/generated/inventory/v1"
	"github.com/mike_jacks/pizza_co/servers/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.InventoryServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	inventory_v1_pb.RegisterInventoryServiceServer(grpcServer, &inventoryServer{})
	reflection.Register(grpcServer)

	log.Printf("Inventory servers is running on port :%d...", config.InventoryServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
