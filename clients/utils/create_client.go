package utils

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateClient initializes and returns a gRPC client for the given service
func CreateClient[T any](host string, port int, clientFactory func(conn *grpc.ClientConn) T) (T, error) {
	address := host + ":" + fmt.Sprintf("%d", port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		var zero T
		return zero, err
	}
	return clientFactory(conn), nil

}
