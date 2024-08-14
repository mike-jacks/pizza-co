package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/mike_jacks/pizza_co/config"
	"github.com/mike_jacks/pizza_co/inventory_service/adapters"
	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
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

func NewServer(inventoryServer inventory_v1_pb.InventoryServiceServer) *Server {
	grpcServer := grpc.NewServer()
	inventory_v1_pb.RegisterInventoryServiceServer(grpcServer, inventoryServer)
	reflection.Register(grpcServer)
	return &Server{
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.InventoryServerPort))
	if err != nil {
		return err
	}
	log.Printf("Inventory servers is running on port :%d...", config.InventoryServerPort)
	os.Stdout.Sync()
	return s.grpcServer.Serve(lis)
}
