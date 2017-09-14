package dao

import (
    "time"
    "github.com/satori/go.uuid"
)

type Purchase struct{
    Items         []InventoryItem 
    TotalCost     float64
    OscCardNumber uint64
    Address       string
    Username      string
    CheckoutDate  time.Time
    Id            *uuid.UUID
}
