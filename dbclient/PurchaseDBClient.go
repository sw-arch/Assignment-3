package dbclient

import (
	"Assignment-3/dao"
	"database/sql"
	"encoding/json"
	"time"

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
			`id Text PRIMARY KEY,
            checkoutdate Text,
            username Text,
            address Text,
            oscnum Integer,
            total Real,
            cart Blob`)
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

func (client PurchaseDBClient) AddPurchase(purchase *dao.Purchase) bool {
	encodedCart, marErr := json.Marshal(purchase.Cart)
	checkErr(marErr)
	statement, prepErr := client.db.Prepare("INSERT INTO purchase (id, checkoutdate, username, address, oscnum, total, cart) VALUES (?, ?, ?, ?, ?, ?, ?);")
	checkErr(prepErr)

	_, err := statement.Exec(purchase.Id, purchase.CheckoutDate.Format(time.RFC3339), purchase.Username, purchase.Address, purchase.OscCardNumber, purchase.TotalCost, encodedCart)
	checkErr(err)
	return true
}

func makePurchasesFromRows(rows *sql.Rows) []dao.Purchase {
	var purchases []dao.Purchase
	for rows.Next() {
		purchase := dao.Purchase{}
		var encodedCart []byte
		var date string
		rowErr := rows.Scan(&purchase.Id, &date, &purchase.Username, &purchase.Address, &purchase.OscCardNumber, &purchase.TotalCost, &encodedCart)
		checkErr(rowErr)

		var cart dao.Cart
		marErr := json.Unmarshal(encodedCart, &cart)
		checkErr(marErr)
		purchase.Cart = &cart
		purchase.CheckoutDate, _ = time.Parse(time.RFC3339, date)

		purchases = append(purchases, purchase)
	}

	return purchases
}
