package utils

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateClient initializes and returns a gRPC client for the given service
func CreateClient[T any](host string, port int, clientFactory func(conn *grpc.ClientConn) T) (T, error) {
	address := host + ":" + fmt.Sprintf("%d", port)
	var conn *grpc.ClientConn
	var err error

	maxRetries := 5
	backoff := 500 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		conn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			return clientFactory(conn), nil
		}

		log.Printf("Failed to connect to %s, retrying in %v... (%d/%d)\n", address, backoff, i+1, maxRetries)
		time.Sleep(backoff)
		backoff *= 2
	}

	var zero T
	return zero, fmt.Errorf("failed to connect to %s after %d retries: %w", address, maxRetries, err)
}
