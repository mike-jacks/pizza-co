// inventory_service/adapters/db/inventory_repository.go

package persistence

import (
	"fmt"
	"strings"

	"github.com/mike_jacks/pizza_co/inventory_service/ports/repository"

	"github.com/mike_jacks/pizza_co/inventory_service/domain/entities"
	"github.com/mike_jacks/pizza_co/inventory_service/domain/types"
	"gorm.io/gorm"
)

type GormInventoryRepository struct {
	repository.InventoryRepository
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *GormInventoryRepository {
	return &GormInventoryRepository{db: db}
}

func (r *GormInventoryRepository) CheckInventory(pizzas []types.Pizza) error {
	// Initialize counters for toppings and crust sizes
	toppingQuantities := make(map[types.Topping]int)
	crustSizeQuantities := make(map[string]int) // Keyed by "CrustType-Size"

	// Accumulate required quantities
	for _, pizza := range pizzas {
		// Accumulate topping quantities
		for _, topping := range pizza.Toppings {
			toppingQuantities[topping] += pizza.Quantity
		}

		// Accumulate crust size quantities (composite key: "CrustType-Size")
		crustKey := fmt.Sprintf("%s-%s", pizza.CrustType, pizza.Size)
		crustSizeQuantities[crustKey] += pizza.Quantity
	}

	// Validate topping inventory
	for topping, requiredQuantity := range toppingQuantities {
		var dbTopping entities.Topping
		err := r.db.Where("name = ?", topping).First(&dbTopping).Error
		if err != nil {
			return fmt.Errorf("topping %s not found in inventory", topping)
		}
		if dbTopping.Quantity < requiredQuantity {
			return fmt.Errorf("not enough quantity for topping %s: required %d, available %d", topping, requiredQuantity, dbTopping.Quantity)
		}
	}

	// Validate crust size inventory
	for crustKey, requiredQuantity := range crustSizeQuantities {
		var dbCrustSize entities.CrustSize
		crustType, crustSize := splitCrustKey(crustKey) // Split the composite key back into CrustType and Size

		// Select the CrustSize entry
		err := r.db.Joins("JOIN crust_type ON crust_type.id = crust_size.crust_type_id").
			Where("crust_type.type = ? AND crust_size.size = ?", crustType, crustSize).
			First(&dbCrustSize).Error
		if err != nil {
			return fmt.Errorf("crust type %s with size %s not found in inventory", crustType, crustSize)
		}
		if dbCrustSize.Quantity < requiredQuantity {
			return fmt.Errorf("not enough quantity for crust type %s with size %s: required %d, available %d", crustType, crustSize, requiredQuantity, dbCrustSize.Quantity)
		}
	}

	// If all quantities are validated, subtract them from the inventory

	// Subtract topping quantities
	for topping, quantity := range toppingQuantities {
		err := r.db.Model(&entities.Topping{}).
			Where("name = ?", topping).
			Update("quantity", gorm.Expr("quantity - ?", quantity)).Error
		if err != nil {
			return fmt.Errorf("failed to update quantity for topping %s: %v", topping, err)
		}
	}

	// Subtract crust size quantities
	for crustKey, quantity := range crustSizeQuantities {
		crustType, crustSize := splitCrustKey(crustKey)

		// Select the CrustSize entry
		var dbCrustSize entities.CrustSize
		err := r.db.Joins("JOIN crust_type ON crust_type.id = crust_size.crust_type_id").
			Where("crust_type.type = ? AND crust_size.size = ?", crustType, crustSize).
			First(&dbCrustSize).Error
		if err != nil {
			return fmt.Errorf("failed to find crust type %s with size %s: %v", crustType, crustSize, err)
		}

		// Now update the quantity
		err = r.db.Model(&dbCrustSize).
			Update("quantity", gorm.Expr("quantity - ?", quantity)).Error
		if err != nil {
			return fmt.Errorf("failed to update quantity for crust type %s with size %s: %v", crustType, crustSize, err)
		}
	}

	return nil
}

// Helper function to split the composite key into CrustType and Size
func splitCrustKey(crustKey string) (crustType, crustSize string) {
	parts := strings.SplitN(crustKey, "-", 2)
	if len(parts) == 2 {
		crustType, crustSize = parts[0], parts[1]
	}
	return crustType, crustSize
}

func (r *GormInventoryRepository) ResetInventory() error {
	return fmt.Errorf("")
}
