package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	common_v1_pb "github.com/mike_jacks/pizza_co/common/ports/grpc/v1"
	inventory_v1_pb "github.com/mike_jacks/pizza_co/inventory_service/ports/grpc/v1"
)

type inventoryServer struct {
	inventory_v1_pb.UnimplementedInventoryServiceServer
}

func (s *inventoryServer) CheckInventory(ctx context.Context, req *inventory_v1_pb.InventoryCheckRequest) (*inventory_v1_pb.InventoryCheckResponse, error) {
	log.Println("Checking Inventory...")
	os.Stdout.Sync()
	time.Sleep(3 * time.Second)
	var pizzas []*common_v1_pb.Pizza = req.GetPizzas()
	var toppings []common_v1_pb.Topping
	var crust_types []common_v1_pb.CrustType
	var crust_sizes []common_v1_pb.Size
	for _, pizza := range pizzas {
		quantity := pizza.GetQuantity()
		for i := 0; i < int(quantity); i++ {
			toppings = append(toppings, pizza.GetToppings()...)
			crust_types = append(crust_types, pizza.GetCrustType())
			crust_sizes = append(crust_sizes, pizza.GetSize())
		}
	}
	toppings_count := make(map[common_v1_pb.Topping]int)
	for _, topping := range toppings {
		toppings_count[topping]++
	}
	crust_types_count := make(map[common_v1_pb.CrustType]int)
	for _, crust_type := range crust_types {
		crust_types_count[crust_type]++
	}
	crust_sizes_count := make(map[common_v1_pb.Size]int)
	for _, crust_size := range crust_sizes {
		crust_sizes_count[crust_size]++
	}

	log.Printf("Order requesting toppings: %v", toppings)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Printf("Order requesting crust types: %v", crust_types)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Printf("Order requesting crust sizes: %v", crust_sizes)
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)
	log.Println("Inventory Check Complete!")
	os.Stdout.Sync()
	time.Sleep(1 * time.Second)

	message := "Your order includes the following toppings:\n"

	for topping, count := range toppings_count {
		message = message + fmt.Sprintf("%dx %v, ", count, topping)
	}
	message = message[:len(message)-1] + ".\n\nYour order includes the following crust types:\n"
	for crust_type, count := range crust_types_count {
		message = message + fmt.Sprintf("%dx %v,", count, crust_type)
	}
	message = message[:len(message)-1] + ".\n\nYour order includes the following crust sizes:\n"
	for crust_size, count := range crust_sizes_count {
		message = message + fmt.Sprintf("%dx %v,", count, crust_size)
	}
	message = message[:len(message)-1] + ". All items in your order are in inventory and available to order!\n"

	return &inventory_v1_pb.InventoryCheckResponse{
		Message:     message,
		ErrorCode:   0,
		IsAvailable: true,
	}, nil

}

func (s *inventoryServer) UpdateInventory(ctx context.Context, req *inventory_v1_pb.UpdateInventoryRequest) (*inventory_v1_pb.UpdateInventoryResponse, error) {
	go func() {
		log.Println("Updating...")
		os.Stdout.Sync()
	}()
	time.Sleep(4 * time.Second)
	go func() {
		log.Println("Inventory updated successfully!")
		os.Stdout.Sync()
	}()
	return &inventory_v1_pb.UpdateInventoryResponse{
		Message:   "Inventory has been updated successfully",
		ErrorCode: 0,
	}, nil

}
