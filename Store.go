package main

import (
    "Assignment-3/dao"
    // "Assignment-3/dbclient"
)

type Store struct {
    // inventoryClient dbclient.InventoryDBClient
}

func (store Store) search(query string) []dao.InventoryItem {
    return []dao.InventoryItem{}
}
