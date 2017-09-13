package main

import "Assignment-3/src/dao"

type Cart struct {
    items     []dao.InventoryItem
    totalCost uint64
}
