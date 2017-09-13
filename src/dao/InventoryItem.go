package dao

import "github.com/satori/go.uuid"

type InventoryItem struct {
    quantityAvailable int
    name              string
    description       string
    price             float64
    id                *uuid.UUID
}
