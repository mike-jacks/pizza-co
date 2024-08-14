package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/joho/godotenv"
	"github.com/mike_jacks/pizza_co/clients/order_management"
	common_v1_pb "github.com/mike_jacks/pizza_co/common/ports/grpc/v1"

	order_management_service_v1_pb "github.com/mike_jacks/pizza_co/order_management_service/ports/grpc/v1"
)

var mu sync.Mutex

func orderRequest(index uint32) *order_management_service_v1_pb.OrderRequest {
	mu.Lock()
	uuid := faker.UUIDDigit()
	firstName := faker.FirstName()
	lastName := faker.LastName()
	phoneNumber := faker.Phonenumber()
	mu.Unlock()
	req := &order_management_service_v1_pb.OrderRequest{
		CustomerInfo: &common_v1_pb.Customer{
			Id:           uuid,
			FirstName:    firstName,
			LastName:     lastName,
			EmailAddress: "fakeemail@email.com",
			DeliveryAddress: &common_v1_pb.Address{
				HouseNumber: "22",
				StreetName:  "Wildcat",
				AptNumber:   "",
				City:        "St George",
				State:       "Utah",
				ZipCode:     "84790",
			},
			PhoneNumber: &common_v1_pb.PhoneNumber{
				Number: phoneNumber,
				Type:   common_v1_pb.PhoneType_MOBILE,
			},
		},
		Pizzas: []*common_v1_pb.Pizza{
			{
				Toppings:     []common_v1_pb.Topping{common_v1_pb.Topping_PEPPERONI, common_v1_pb.Topping_BLACK_OLIVES, common_v1_pb.Topping_ANCHOVIES},
				Size:         common_v1_pb.Size_EXTRA_LARGE,
				CrustType:    common_v1_pb.CrustType_NEW_YORK,
				ExtraOptions: []common_v1_pb.Extra{},
				Quantity:     index,
			},
		},
		PaymentMethod: &common_v1_pb.Payment{
			PaymentType:      common_v1_pb.PaymentType_CREDIT_CARD,
			PaymentTimeframe: common_v1_pb.PaymentTimeframe_PREPAID,
			TotalOrderAmount: fmt.Sprint(13 * index),
		},
	}
	return req
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Number of concurrent requests
	numRequests := 50

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 1; i <= numRequests; i++ {
		orderManagementServerHost := os.Getenv("ORDER_MANAGEMENT_SERVICE_HOST")
		var port int
		if i%2 == 0 {
			port = 9000
		} else {
			port = 9000
		}

		orderManagementClient, err := order_management.CreateOrderManagementClient(orderManagementServerHost, port)
		if err != nil {
			log.Fatalf("Failed to create order manaagement client ; %v", err)
			os.Stdout.Sync()
		}
		log.Printf("Making request %d to %s:%d", i, orderManagementServerHost, port)

		time.Sleep(100 * time.Millisecond)
		go func(i int) {
			defer wg.Done() // Mark this goroutine as done when it finishes
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Request %d panicked: %v", i, r)
				}
			}()
			req := orderRequest(uint32(i))

			// Make the gRPC request
			ctx := context.Background()
			stream, err := orderManagementClient.PlaceOrder(ctx, req)
			if err != nil {
				log.Printf("Request %d failed: %v", i, err)
				return
			}
			// Handle the streaming response
			for {
				response, err := stream.Recv()
				if err == io.EOF {
					break // End of stream
				}
				if err != nil {
					log.Printf("Request %d failed to receive response: %v", i, err)
					return
				}

				// Process each response message
				log.Printf("Request %d received response: %v", i, response)
			}

		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	log.Println("All requests completed.")

	time.Sleep(5 * time.Second)
}
