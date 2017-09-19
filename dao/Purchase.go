package dao

import (
	"time"

	"github.com/satori/go.uuid"
)

type Purchase struct {
	Id            *uuid.UUID
	CheckoutDate  time.Time
	Username      string
	Address       string
	OscCardNumber uint64
	TotalCost     float64
	Cart          Cart
}
