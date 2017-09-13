package main

import "Assignment-3/dao"

type Cart struct {
    items     []dao.InventoryItem
    totalCost uint64
}
