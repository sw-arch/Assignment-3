package dbclient

import (
    "github.com/satori/go.uuid"
    "Assignment-3/dao"
)

type InventoryDBClient struct {
}

func (client InventoryDBClient) GetItemByID(id uuid.UUID) dao.InventoryItem {
    return dao.InventoryItem{}
}

func (client InventoryDBClient) SetItemQuantity(quantity uint64) bool {
    return false
}

func (client InventoryDBClient) Reserve(item dao.InventoryItem, quantity uint64) bool {
    return false
}

func (client InventoryDBClient) Release(item dao.InventoryItem, quantity uint64) bool {
    return false
}

func (client InventoryDBClient) Remove(item dao.InventoryItem, quantity uint64) bool {
    return false
}