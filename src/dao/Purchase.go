package dao

import "github.com/satori/go.uuid"

import "time"

type Purchase struct{
    items         []InventoryItem 
    totalCost     float64
    oscCardNumber uint64
    address       string
    username      string
    checkoutDate  time.Time
    id            *uuid.UUID
}
