package entities

type Topping struct {
	ID       uint            `gorm:"primaryKey"`
	Name     string          `gorm:"type:varchar(100);unique;not null"`
	Quantity int             `gorm:"not null"`
	Items    []InventoryItem `gorm:"many2many:inventory_item_topping"`
}

type CrustType struct {
	ID    uint        `gorm:"primaryKey"`
	Type  string      `gorm:"type:varchar(50);unique;not null"`
	Sizes []CrustSize `gorm:"foreignKey:CrustTypeID"`
}

// CrustSize represents the size of the crust, e.g., small, medium, large, etc.
type CrustSize struct {
	ID          uint            `gorm:"primaryKey"`
	Size        string          `gorm:"type:varchar(50);not null"`
	Quantity    int             `gorm:"not null"` // Track quantity of this specific crust size tied to crust type
	CrustTypeID uint            `gorm:"not null"` // Foreign key to CrustType
	CrustType   CrustType       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items       []InventoryItem `gorm:"foreignKey:CrustSizeID"`
}

// InventoryItem represents a specific item in the inventory with a given crust type and size, and its associated toppings.
type InventoryItem struct {
	ID          uint      `gorm:"primaryKey"`
	Toppings    []Topping `gorm:"many2many:inventory_item_toppings;"`
	CrustSizeID uint      `gorm:"not null"`
	CrustSize   CrustSize `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity    int       `gorm:"not null"`
}

// InventoryItemTopping represents the join table for the many-to-many relationship between InventoryItem and Topping.
type InventoryItemTopping struct {
	InventoryItemID uint `gorm:"primaryKey"`
	ToppingID       uint `gorm:"primaryKey"`
}
