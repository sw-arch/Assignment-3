package dbclient

import (
	"Assignment-3/dao"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

type InventoryDBClient struct {
	database *sql.DB
}

var inventoryDBInstance *InventoryDBClient

func GetInventoryDBClient() *InventoryDBClient {
	if inventoryDBInstance == nil {
		inventoryDBInstance = &InventoryDBClient{initializeInventoryDB()}
		inventoryDBInstance.createInventoryTable()
	}
	return inventoryDBInstance
}

func initializeInventoryDB() *sql.DB {
	db, err := sql.Open("sqlite3", "inventory.db")
	checkErr(err)
	if db == nil {
		panic("DB is nil!")
	}

	return db
}

func (client InventoryDBClient) createInventoryTable() {
	checkStatement := "SELECT name FROM sqlite_master WHERE type='table' AND name='inventory';"
	result, err := client.database.Exec(checkStatement)
	checkErr(err)
	if result == nil {
		createStatement := `Create Table inventory (
            inventory_id Text primary key,
            name Text,
            description Text,
            price Real
            );`
		_, err := client.database.Exec(createStatement)
		checkErr(err)
	}
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
