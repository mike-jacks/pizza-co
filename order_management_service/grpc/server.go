package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/mike_jacks/pizza_co/clients/inventory"
	"github.com/mike_jacks/pizza_co/clients/utils"
	"github.com/mike_jacks/pizza_co/config"
	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/order_management_service/ports/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func GetENV() error {
	if os.Getenv("USE_ENV_FILE") != "false" {
		if err := godotenv.Load("../.env"); err != nil {
			return fmt.Errorf("no .env file found or error loading it: %v", err)
		} else {
			log.Println(".env file loaded successfully")
			os.Stdout.Sync()
		}
	} else {
		log.Println("Skipping loading .env file as USE_ENV_FILEs set to false")
		os.Stdout.Sync()
	}
	return nil
}

func getInventoryClient() inventory_v1_pb.InventoryServiceClient {
	inventoryClientHost := os.Getenv("INVENTORY_SERVICE_HOST")
	// Create inventory client
	inventoryClient, err := utils.CreateClient(inventoryClientHost, config.InventoryServerPort, inventory.NewInventoryClient)
	if err != nil {
		log.Fatalf("Failed to create inventory client ; %v", err)
		os.Stdout.Sync()
	}
	return inventoryClient
}

func NewServer() *Server {
	grpcServer := grpc.NewServer()
	orderManagementServer := NewOrderManagementServer(getInventoryClient())
	order_management_v1_pb.RegisterOrderManagementServiceServer(grpcServer, orderManagementServer)
	reflection.Register(grpcServer)
	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.OrderManagementServerPort))
	if err != nil {
		return err
	}
	log.Printf("Server is running on port :%d...", config.OrderManagementServerPort)
	return s.grpcServer.Serve(lis)
}
