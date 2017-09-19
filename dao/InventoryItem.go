package dao

import "github.com/satori/go.uuid"

type InventoryItem struct {
	Id                uuid.UUID
	Name              string
	Description       string
	Category          string
	Price             float64
	QuantityAvailable uint64
}
