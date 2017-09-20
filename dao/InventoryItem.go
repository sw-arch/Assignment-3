package dao

import "github.com/satori/go.uuid"

type InventoryItem struct {
	Id                uuid.UUID
	Name              string
	Description       string
	Category          string
	AttributeOne      string
	AttributeTwo      string
	Price             float64
	QuantityAvailable int64
}
