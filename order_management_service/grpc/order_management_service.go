package grpc

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
	order_management_v1_pb "github.com/mike_jacks/pizza_co/order_management_service/ports/grpc/v1"
	"google.golang.org/grpc/status"
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

func checkInventory(s *orderManagementServer, req *inventory_v1_pb.InventoryCheckRequest) (*inventory_v1_pb.InventoryCheckResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resp, err := s.inventoryClient.CheckInventory(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		return resp, fmt.Errorf(st.Message())
	}
	return resp, nil
}

func (s *orderManagementServer) PlaceOrder(req *order_management_v1_pb.OrderRequest, stream order_management_v1_pb.OrderManagementService_PlaceOrderServer) error {

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
			log.Printf("Received order for customer: %v %v", req.GetCustomerInfo().GetFirstName(), req.GetCustomerInfo().GetLastName())
			os.Stdout.Sync()
			time.Sleep(2 * time.Second)
			message = "Your order has been received"
		case "CHECKING_INVENTORY":
			log.Printf("Checking inventory")
			os.Stdout.Sync()
			time.Sleep(2 * time.Second)
			res := &order_management_v1_pb.OrderResponse{
				OrderId: orderID,
				Status:  stage,
				Message: "Currently Checking Inventory....standby...",
			}
			if err := stream.Send(res); err != nil {
				log.Println(err)
				os.Stdout.Sync()
				return err
			}
			pizzas := req.GetPizzas()
			req := &inventory_v1_pb.InventoryCheckRequest{
				Pizzas: pizzas,
			}

			resp, err := checkInventory(s, req)
			if err != nil {
				message = err.Error()
				log.Printf("Error checking inventory: %v", err)
				os.Stdout.Sync()
				res := &order_management_v1_pb.OrderResponse{
					OrderId: orderID,
					Status:  stage,
					Message: message,
				}
				if stream_err := stream.Send(res); stream_err != nil {
					log.Println(err)
					os.Stdout.Sync()
					return stream_err
				}
				return nil
			}
			time.Sleep(2 * time.Second)
			message = fmt.Sprintf("%v", resp.GetMessage())

		case "PROCESSING_PAYMENT":
			log.Printf("Begin Processing Payment")
			os.Stdout.Sync()
			time.Sleep(2 * time.Second)
			message = "Your payment successfully went through. Thank you!"
		case "PROCESSING_ORDER":
			log.Print("Begin processiong order.")
			os.Stdout.Sync()
			message = "Order Processing complete!"
		case "COMPLETE":
			log.Print("Order Process complete, sending final message")
			os.Stdout.Sync()
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
			os.Stdout.Sync()
			return err
		}

		// Simulate delay in processing each stage
		time.Sleep(2 * time.Second)
	}
	return nil
}
