// inventory_service/config/db.go

package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mike_jacks/pizza_co/inventory_service/domain/entities"
	"github.com/mike_jacks/pizza_co/inventory_service/domain/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {

	// Define the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT"))

	fmt.Println("DSN:", dsn) // Print the DSN for debugging

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

	// Retrieve the underlying *sql.DB object to configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	// Set the maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(25)

	// Set the maximum number of idle connections in the pool
	sqlDB.SetMaxIdleConns(25)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	// Auto-migrate the schema
	err = db.AutoMigrate(&entities.InventoryItem{}, &entities.Topping{}, &entities.InventoryItemTopping{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	initializeData(db)
	log.Println("Database has been reset with default data")

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
