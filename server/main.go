package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"time"

	pb "github.com/mike-jacks/pizza-co/generated/order_management/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type orderManagementServer struct {
	pb.UnimplementedOrderManagementServiceServer
}

func generateOrderNumber() string {
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	randomNum := rng.IntN(900000) + 100000
	orderNumber := fmt.Sprintf("ORDER24%d", randomNum)
	return orderNumber
}

func (s *orderManagementServer) PlaceOrder(req *pb.OrderRequest, stream pb.OrderManagementService_PlaceOrderServer) error {
	log.Printf("Received order for customer: %v %v", req.GetCustomerInfo().GetFirstName(), req.GetCustomerInfo().GetLastName())

	stages := []pb.Status{
		pb.Status_RECEIVED,
		pb.Status_CHECKING_INVENTORY,
		pb.Status_PROCESSING_PAYMENT,
		pb.Status_PROCESSING_ORDER,
		pb.Status_COMPLETE,
	}

	orderID := generateOrderNumber()

	for _, stage := range stages {
		var message string

		switch stage.String() {
		case "RECEIVED":
			message = "Your order has been received"
		case "CHECKING_INVENTORY":
			message = "Checking inventory to see if items are in stock for your oder"
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

		res := &pb.OrderResponse{
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

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderManagementServiceServer(grpcServer, &orderManagementServer{})
	reflection.Register(grpcServer)

	log.Printf("Server is running on port :%d...", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
