package dao

import "github.com/satori/go.uuid"

type InventoryItem struct {
	Id                *uuid.UUID
	Name              string
	Description       string
	Price             float64
	QuantityAvailable int
}
