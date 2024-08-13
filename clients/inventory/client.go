package inventory

import (
	"fmt"

	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewInventoryClient creates a new InventoryServiceClient.
func NewInventoryClient(conn *grpc.ClientConn) inventory_v1_pb.InventoryServiceClient {
	return inventory_v1_pb.NewInventoryServiceClient(conn)
}

// CreateInventoryClient initializes and returns an InventoryServiceClient.
func CreateInventoryClient(host string, port int) (inventory_v1_pb.InventoryServiceClient, error) {
	address := host + ":" + fmt.Sprintf("%d", port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewInventoryClient(conn), nil
}
