package repository

type InventoryRepository interface {
	CheckInventory(item any) error
	UpdateInventory() error
}
