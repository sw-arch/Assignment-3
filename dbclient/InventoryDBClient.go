package dbclient

import (
    "Assignment-3/dao"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
    "github.com/satori/go.uuid"
)

type InventoryDBClient struct {
    db *sql.DB
}

var inventoryDBInstance *InventoryDBClient

func GetInventoryDBClient() *InventoryDBClient {
    if inventoryDBInstance == nil {
        inventoryDBInstance = &InventoryDBClient{initializeDB("inventory.db")}
        createTable(inventoryDBInstance.db, "inventory",
            `inventory_id Integer primary key,
            name Text,
            description Text,
            price Real,
            quantity_on_hand Integer,
            quantity_reserved Integer`)
    }
    return inventoryDBInstance
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
