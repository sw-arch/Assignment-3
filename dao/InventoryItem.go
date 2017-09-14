package dao

import "github.com/satori/go.uuid"

type InventoryItem struct {
    QuantityAvailable int
    Name              string
    Description       string
    Price             float64
    Id                *uuid.UUID
}
