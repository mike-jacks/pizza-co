package adapters

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	common_v1_pb "github.com/mike_jacks/pizza_co/common/ports/grpc/v1"
	"github.com/mike_jacks/pizza_co/inventory_service/adapters/persistence"
	"github.com/mike_jacks/pizza_co/inventory_service/domain/actions"
	"github.com/mike_jacks/pizza_co/inventory_service/domain/types"
	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
)

type InventoryServer struct {
	inventory_v1_pb.UnimplementedInventoryServiceServer
	repo *persistence.GormInventoryRepository
}

func NewInventoryServer(repo *persistence.GormInventoryRepository) *InventoryServer {
	return &InventoryServer{repo: repo}
}

func convertToppings(toppings []common_v1_pb.Topping) []types.Topping {
	var result []types.Topping
	for _, topping := range toppings {
		result = append(result, types.Topping(topping.String()))
	}
	return result
}

func (s *InventoryServer) CheckInventory(ctx context.Context, req *inventory_v1_pb.InventoryCheckRequest) (*inventory_v1_pb.InventoryCheckResponse, error) {
	log.Println("Checking Inventory...")
	time.Sleep(3 * time.Second)

	// Convert gRPC request to domain types
	pizzas := []types.Pizza{}
	toppings := []types.Topping{}
	crustTypes := []types.CrustType{}
	crustSizes := []types.Size{}
	for _, pizza := range req.GetPizzas() {
		pizzas = append(pizzas, types.Pizza{
			Toppings:  convertToppings(pizza.GetToppings()),
			CrustType: types.CrustType(pizza.GetCrustType().String()),
			Size:      types.Size(pizza.GetSize().String()),
			Quantity:  int(pizza.GetQuantity()),
		})
		toppings = append(toppings, convertToppings(pizza.GetToppings())...)
		crustTypes = append(crustTypes, types.CrustType(pizza.GetCrustType().String()))
		crustSizes = append(crustSizes, types.Size(pizza.GetSize().String()))

	}

	// Call the domain logic
	toppingsCount, crustTypesCount, crustSizesCount := actions.CheckInventory(pizzas)

	log.Printf("Order requesting toppings: %v", toppings)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Printf("Order requesting crust types: %v", crustTypes)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Printf("Order requesting crust sizes: %v", crustSizes)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Println("Inventory Check Complete!")
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)

	message := "Your order includes the following toppings:\n"

	for topping, count := range toppingsCount {
		message = message + fmt.Sprintf("%dx %v, ", count, topping)
	}
	message = message[:len(message)-1] + ".\n\nYour order includes the following crust types:\n"
	for crust_type, count := range crustTypesCount {
		message = message + fmt.Sprintf("%dx %v,", count, crust_type)
	}
	message = message[:len(message)-1] + ".\n\nYour order includes the following crust sizes:\n"
	for crust_size, count := range crustSizesCount {
		message = message + fmt.Sprintf("%dx %v,", count, crust_size)
	}
	message = message[:len(message)-1] + ". All items in your order are in inventory and available to order!\n"

	if err := s.repo.CheckInventory(pizzas); err != nil {
		message = fmt.Sprintf("Error: %v", err.Error())
		return &inventory_v1_pb.InventoryCheckResponse{
			Message:     message,
			ErrorCode:   0,
			IsAvailable: false,
		}, err
	}

	// Build gRPC response
	return &inventory_v1_pb.InventoryCheckResponse{
		Message:     message,
		ErrorCode:   0,
		IsAvailable: true,
	}, nil
}

func (s *InventoryServer) UpdateInventory(ctx context.Context, req *inventory_v1_pb.UpdateInventoryRequest) (*inventory_v1_pb.UpdateInventoryResponse, error) {
	log.Println("Updating...")
	os.Stdout.Sync()
	time.Sleep(4 * time.Second)
	log.Println("Inventory updated successfully!")
	os.Stdout.Sync()
	return &inventory_v1_pb.UpdateInventoryResponse{
		Message:   "Inventory has been updated successfully",
		ErrorCode: 0,
	}, nil
}
