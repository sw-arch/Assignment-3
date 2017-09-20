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
		createTable(inventoryDBInstance.db, "categories",
			`identifier Text PRIMARY KEY,
			description Text`)
		createTable(inventoryDBInstance.db, "inventory",
			`inventory_id Text PRIMARY KEY,
            name Text,
			description Text,
			category Text,
            price Real,
            quantity_on_hand Integer,
			quantity_reserved Integer,
			FOREIGN KEY (category) REFERENCES categories(identifier)`)
	}
	return inventoryDBInstance
}

func (client InventoryDBClient) GetAllItems() []dao.InventoryItem {
	rows, err := client.db.Query("SELECT * FROM inventory;")
	checkErr(err)
	defer rows.Close()

	return makeInventoryItemsFromRows(rows)
}

func (client InventoryDBClient) GetItemsByCategory(category string) []dao.InventoryItem {
	statement, prepErr := client.db.Prepare("SELECT * FROM inventory WHERE category=?")
	checkErr(prepErr)

	rows, queryErr := statement.Query(category)
	checkErr(queryErr)
	defer rows.Close()

	return makeInventoryItemsFromRows(rows)
}

func (client InventoryDBClient) GetItemByID(id uuid.UUID) dao.InventoryItem {
	statement, prepErr := client.db.Prepare("SELECT * FROM inventory WHERE inventory_id=?;")
	checkErr(prepErr)

	row := statement.QueryRow(id)
	return makeInventoryItemFromRow(row)
}

func (client InventoryDBClient) SetItemQuantity(itemId uuid.UUID, quantity int64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_on_hand=? where inventory_id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, itemId.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Reserve(item dao.InventoryItem, quantity int64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_reserved=quantity_reserved + ? where inventory_id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Release(item dao.InventoryItem, quantity int64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_reserved=quantity_reserved - ? where inventory_id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) Remove(item dao.InventoryItem, quantity int64) bool {
	statement, prepErr := client.db.Prepare("UPDATE inventory SET quantity_on_hand=quantity_on_hand - ?, quantity_reserved=quantity_reserved - ? where inventory_id=?;")
	checkErr(prepErr)

	res, execErr := statement.Exec(quantity, quantity, item.Id.String())
	checkErr(execErr)

	rowCount, err := res.RowsAffected()
	checkErr(err)

	return rowCount != 0
}

func (client InventoryDBClient) GetAvailableCategories() []string {
	var categories []string

	rows, err := client.db.Query("SELECT DISTINCT category FROM inventory")
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

func (client InventoryDBClient) GetCategoryDescriptions() [][]string {
	var categories [][]string

	rows, err := client.db.Query("SELECT identifier, description FROM categories")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var identifier string
		var description string
		err = rows.Scan(&identifier, &description)
		checkErr(err)
		categories = append(categories, []string{identifier, description})
	}

	return categories
}

func makeInventoryItemFromRow(row *sql.Row) dao.InventoryItem {
	item := dao.InventoryItem{}
	var quantityOnHand int64
	var quantityReserved int64
	rowErr := row.Scan(&item.Id, &item.Name, &item.Description, &item.Category, &item.Price, &quantityOnHand, &quantityReserved)
	checkErr(rowErr)
	item.QuantityAvailable = quantityOnHand
	return item
}

func makeInventoryItemsFromRows(rows *sql.Rows) []dao.InventoryItem {
	var items []dao.InventoryItem
	for rows.Next() {
		item := dao.InventoryItem{}
		var quantityOnHand int64
		var quantityReserved int64
		rowErr := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Category, &item.Price, &quantityOnHand, &quantityReserved)
		checkErr(rowErr)
		item.QuantityAvailable = quantityOnHand
		items = append(items, item)
	}

	return items
}
