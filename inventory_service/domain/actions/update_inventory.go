package actions

import (
	"log"

	"github.com/mike_jacks/pizza_co/inventory_service/domain/entities"
)

func UpdateInventory(items []entities.InventoryItem) {
	// Implement inventory update logic here
	for _, item := range items {
		log.Printf("Updating inventory for item ide: %d - %d units", item.ID, item.Quantity)
		// Update logic here, possibly interacting with a database
	}
}
