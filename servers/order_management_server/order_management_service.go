package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	inventory_v1_pb "github.com/mike_jacks/pizza_co/generated/inventory/v1"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/generated/order_management/v1"
)

type orderManagementServer struct {
	order_management_v1_pb.UnimplementedOrderManagementServiceServer
	inventoryClient inventory_v1_pb.InventoryServiceClient
}

func NewOrderManagementServer(inventoryClient inventory_v1_pb.InventoryServiceClient) *orderManagementServer {
	return &orderManagementServer{
		inventoryClient: inventoryClient,
	}
}

func generateOrderNumber() string {
	// Create a new random number generator with a unique seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 100000 and 999999
	randomNum := rng.Intn(900000) + 100000
	orderNumber := fmt.Sprintf("ORDER24%d", randomNum)
	return orderNumber
}

func (s *orderManagementServer) PlaceOrder(req *order_management_v1_pb.OrderRequest, stream order_management_v1_pb.OrderManagementService_PlaceOrderServer) error {
	log.Printf("Received order for customer: %v %v", req.GetCustomerInfo().GetFirstName(), req.GetCustomerInfo().GetLastName())

	stages := []order_management_v1_pb.Status{
		order_management_v1_pb.Status_RECEIVED,
		order_management_v1_pb.Status_CHECKING_INVENTORY,
		order_management_v1_pb.Status_PROCESSING_PAYMENT,
		order_management_v1_pb.Status_PROCESSING_ORDER,
		order_management_v1_pb.Status_COMPLETE,
	}

	orderID := generateOrderNumber()

	for _, stage := range stages {
		var message string

		switch stage.String() {
		case "RECEIVED":
			message = "Your order has been received"
		case "CHECKING_INVENTORY":
			pizzas := req.GetPizzas()
			req := &inventory_v1_pb.InventoryCheckRequest{
				Pizzas: pizzas,
			}
			resp, err := s.inventoryClient.CheckInventory(context.Background(), req)
			if err != nil {
				log.Fatalf("Error checking inventory: %v", err)
			}
			message = fmt.Sprintf("%v", resp.GetMessage())

		case "PROCESSING_PAYMENT":
			message = "Items in stock! Processing payment for your order"
		case "PROCESSING_ORDER":
			message = "Payment complete! Processing your order"
		case "COMPLETE":
			switch req.GetPaymentMethod().GetPaymentTimeframe().String() {
			case "PREPAID":
				message = "Order is ready for pickup or delivery."
			case "PAYLATER":
				message = "Order ready for pickup and payment"
			default:
				message = "Payment Timeframe not specified. Come on in and pickup"
			}
		default:
			message = "Error with processing your order"
		}

		res := &order_management_v1_pb.OrderResponse{
			OrderId: orderID,
			Status:  stage,
			Message: message,
		}

		if err := stream.Send(res); err != nil {
			log.Println(err)
		}

		// Simulate delay in processing each stage
		time.Sleep(4 * time.Second)
	}
	return nil
}
