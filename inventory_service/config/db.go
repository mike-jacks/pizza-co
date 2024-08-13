// inventory_service/config/db.go

package config

import (
	"log"

	"github.com/mike_jacks/pizza_co/inventory_service/domain/entities"
	"github.com/mike_jacks/pizza_co/inventory_service/domain/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	// Define the DSN (Data Source Name)
	dsn := "host=localhost user=postgres password=postgres dbname=pizza port=5435 sslmode=disable"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Drop existing tables if they exist
	err = db.Migrator().DropTable(&entities.InventoryItem{}, &entities.Topping{}, &entities.InventoryItemTopping{}, entities.CrustType{}, entities.CrustSize{})
	if err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&entities.InventoryItem{}, &entities.Topping{}, &entities.InventoryItemTopping{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	initializeData(db)

	return db
}

func initializeData(db *gorm.DB) {
	toppings := []types.Topping{
		types.PEPPERONI, types.MUSHROOMS, types.ONIONS,
		types.SAUSAGE, types.BACON, types.BLACK_OLIVES,
		types.GREEN_PEPPERS, types.PINEAPPLE, types.ANCHOVIES,
	}

	crustTypes := []types.CrustType{
		types.THIN, types.REGULAR, types.STUFFED,
		types.NEW_YORK, types.DEEP_DISH, types.GLUTEN_FREE,
	}

	sizes := []types.Size{
		types.SMALL, types.MEDIUM, types.LARGE, types.EXTRA_LARGE,
	}

	// Insert toppings
	for _, topping := range toppings {
		db.Create(&entities.Topping{
			Name:     string(topping),
			Quantity: 100,
		})
	}

	// Insert crust types and sizes
	for _, crustType := range crustTypes {
		// Create each crust type
		crust := &entities.CrustType{
			Type: string(crustType),
		}
		db.Create(crust)

		// For each crust type, create associated sizes
		for _, size := range sizes {
			db.Create(&entities.CrustSize{
				Size:        string(size),
				Quantity:    100,
				CrustTypeID: crust.ID,
				CrustType:   *crust,
			})
		}
	}
}
