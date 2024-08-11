package order_management

import (
	"fmt"

	order_management_v1_pb "github.com/mike_jacks/pizza_co/generated/order_management/v1"
	"google.golang.org/grpc"
)

// NewOrderManagementClient creates a new OrderManagementServiceClient.
func NewOrderManagementClient(conn *grpc.ClientConn) order_management_v1_pb.OrderManagementServiceClient {
	return order_management_v1_pb.NewOrderManagementServiceClient(conn)
}

// CreateOrderManagementClient initializes and returns an OrderManagerClient.
func CreateOrderManagementClient(host string, port int) (order_management_v1_pb.OrderManagementServiceClient, error) {
	address := host + ":" + fmt.Sprintf("%d", port)
	conn, err := grpc.NewClient(address)
	if err != nil {
		return nil, err
	}
	return NewOrderManagementClient(conn), nil
}
