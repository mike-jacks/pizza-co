package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mike_jacks/pizza_co/config"
	"github.com/mike_jacks/pizza_co/inventory_service/adapters"
	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	grpcServer := grpc.NewServer()
	inventory_v1_pb.RegisterInventoryServiceServer(grpcServer, &adapters.InventoryServer{})
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
