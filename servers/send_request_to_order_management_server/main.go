package main

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/joho/godotenv"
	"github.com/mike_jacks/pizza_co/clients/order_management"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/generated/order_management/v1"
)

var mu sync.Mutex

func orderRequest(index uint32) *order_management_v1_pb.OrderRequest {
	mu.Lock()
	uuid := faker.UUIDDigit()
	firstName := faker.FirstName()
	lastName := faker.LastName()
	phoneNumber := faker.Phonenumber()
	mu.Unlock()
	req := &order_management_v1_pb.OrderRequest{
		CustomerInfo: &order_management_v1_pb.Customer{
			Id:           uuid,
			FirstName:    firstName,
			LastName:     lastName,
			EmailAddress: "fakeemail@email.com",
			DeliveryAddress: &order_management_v1_pb.Address{
				HouseNumber: "22",
				StreetName:  "Wildcat",
				AptNumber:   "",
				City:        "St George",
				State:       "Utah",
				ZipCode:     "84790",
			},
			PhoneNumber: &order_management_v1_pb.PhoneNumber{
				Number: phoneNumber,
				Type:   order_management_v1_pb.PhoneType_MOBILE,
			},
		},
		Pizzas: []*order_management_v1_pb.Pizza{
			{
				Toppings:     []order_management_v1_pb.Pizza_Topping{order_management_v1_pb.Pizza_PEPPERONI, order_management_v1_pb.Pizza_BLACK_OLIVES},
				Size:         order_management_v1_pb.Pizza_EXTRA_LARGE,
				CrustType:    order_management_v1_pb.Pizza_NEW_YORK,
				ExtraOptions: []order_management_v1_pb.Pizza_Extra{},
				Quantity:     index,
			},
		},
		PaymentMethod: &order_management_v1_pb.Payment{
			PaymentType:      order_management_v1_pb.PaymentType_CREDIT_CARD,
			PaymentTimeframe: order_management_v1_pb.PaymentTimeframe_PREPAID,
			TotalOrderAmount: string(13 * index),
		},
	}
	return req
}

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Number of concurrent requests
	numRequests := 500

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 1; i <= numRequests; i++ {
		orderManagementServerHost := os.Getenv("ORDER_MANAGEMENT_SERVICE_HOST")
		orderManagementClient, err := order_management.CreateOrderManagementClient(orderManagementServerHost, 9000)
		if err != nil {
			log.Fatalf("Failed to create order manaagement client ; %v", err)
			os.Stdout.Sync()
		}
		log.Printf("Making request %d to %s", i, orderManagementServerHost)

		time.Sleep(10 * time.Millisecond)
		go func(i int) {
			defer wg.Done() // Mark this goroutine as done when it finishes
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Request %d panicked: %v", i, r)
				}
			}()
			req := orderRequest(uint32(i))

			// Make the gRPC request
			stream, err := orderManagementClient.PlaceOrder(context.Background(), req)
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
