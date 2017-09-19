package dbclient

import (
	"Assignment-3/dao"
	"database/sql"
	"fmt"

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
			`inventory_id Text primary key,
            name Text,
			description Text,
			category Text,
            price Real,
            quantity_on_hand Integer,
            quantity_reserved Integer`)
	}
	return inventoryDBInstance
}

func (client InventoryDBClient) GetAllItems() []dao.InventoryItem {
	rows, err := client.db.Query("SELECT * FROM inventory;")
	checkErr(err)
	defer rows.Close()

	var items []dao.InventoryItem
	for rows.Next() {
		newItem := dao.InventoryItem{}
		var quantityReserved int
		rowErr := rows.Scan(&newItem.Id, &newItem.Name, &newItem.Description, &newItem.Category, &newItem.Price, &newItem.QuantityAvailable, &quantityReserved)
		checkErr(rowErr)

		items = append(items, newItem)
	}

	return items
}

func (client InventoryDBClient) GetItemsByCategory(category string) []dao.InventoryItem {
	fmt.Println("getting items in category " + category)
	statement, prepErr := client.db.Prepare("SELECT * FROM inventory WHERE category=?")
	checkErr(prepErr)

	rows, queryErr := statement.Query(category)
	checkErr(queryErr)
	defer rows.Close()

	var items []dao.InventoryItem
	for rows.Next() {
		newItem := dao.InventoryItem{}
		var quantityReserved int
		rowErr := rows.Scan(&newItem.Id, &newItem.Name, &newItem.Description, &newItem.Category, &newItem.Price, &newItem.QuantityAvailable, &quantityReserved)
		checkErr(rowErr)

		items = append(items, newItem)
	}

	return items
}

func (client InventoryDBClient) GetItemByID(id uuid.UUID) dao.InventoryItem {
	statement, prepErr := client.db.Prepare("SELECT * FROM inventory WHERE inventory_id=?;")
	checkErr(prepErr)

	row := statement.QueryRow(id)
	var uUID uuid.UUID
	var name string
	var description string
	var category string
	var price float64
	var quantityOnHand int
	var quantityReserved int
	err := row.Scan(&uUID, &name, &description, &price, &quantityOnHand, &quantityReserved)
	checkErr(err)
	return dao.InventoryItem{id, name, category, description, price, quantityOnHand}
}

func (client InventoryDBClient) SetItemQuantity(itemId uuid.UUID, quantity uint64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_on_hand=? where id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, itemId.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Reserve(item dao.InventoryItem, quantity uint64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_reserved=quantity_reserved - ? where id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Release(item dao.InventoryItem, quantity uint64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_reserved=quantity_reserved + ? where id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Remove(item dao.InventoryItem, quantity uint64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_on_hand=quantity_on_hand - ? where id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) GetAvailableCategories() []string {
	var categories []string

	rows, err := client.db.Query("SELECT category FROM inventory")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		checkErr(err)
		categories = append(categories, category)
	}

	return categories
}
