package dbclient

import (
	"Assignment-3/dao"
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type PurchaseDBClient struct {
	db *sql.DB
}

var purchaseDBInstance *PurchaseDBClient

func GetPurchaseDBClient() *PurchaseDBClient {
	if purchaseDBInstance == nil {
		purchaseDBInstance = &PurchaseDBClient{initializeDB("purchase.db")}
		createTable(purchaseDBInstance.db, "purchase",
			`id Text primary key,
            checkoutdate Numeric,
            username Text,
            address Text,
            oscnum Integer,
            total Real,
            items Blob`)
	}
	return purchaseDBInstance
}

func (client PurchaseDBClient) GetPurchasesByUsername(username string) []dao.Purchase {
	statement, prepErr := client.db.Prepare("SELECT * FROM purchase WHERE username=?;")
	checkErr(prepErr)

	rows, queryErr := statement.Query(username)
	checkErr(queryErr)
	defer rows.Close()

	return makePurchasesFromRows(rows)
}

func (client PurchaseDBClient) AddPurchase(purchase dao.Purchase) bool {
	items, marErr := json.Marshal(purchase.Items)
	checkErr(marErr)
	statement, prepErr := client.db.Prepare("INSERT INTO purchase (id, checkoutdate, username, address, oscnum, total, items) VALUES (?, ?, ?, ?, ?, ?, ?);")
	checkErr(prepErr)

	_, err := statement.Exec(purchase.Id, purchase.CheckoutDate, purchase.Username, purchase.Address, purchase.OscCardNumber, purchase.TotalCost, items)
	checkErr(err)
	return true
}

func makePurchasesFromRows(rows *sql.Rows) []dao.Purchase {
	var purchases []dao.Purchase
	for rows.Next() {
		purchase := dao.Purchase{}
		var encodedItems []byte
		rowErr := rows.Scan(&purchase.Id, &purchase.CheckoutDate, &purchase.Username, &purchase.Address, &purchase.OscCardNumber, &purchase.TotalCost, &encodedItems)
		checkErr(rowErr)

		var items []dao.InventoryItem
		marErr := json.Unmarshal(encodedItems, items)
		checkErr(marErr)
		purchase.Items = items

		purchases = append(purchases, purchase)
	}

	return purchases
}
