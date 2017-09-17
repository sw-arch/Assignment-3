package dbclient

import (
	"Assignment-3/dao"
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
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

func (client PurchaseDBClient) GetPurchaseByID(id uuid.UUID) dao.Purchase {
	statement, prepErr := client.db.Prepare("SELECT * FROM purchase WHERE id=?;")
	checkErr(prepErr)

	row := statement.QueryRow(id)
	var uUID uuid.UUID
	var checkoutDate time.Time
	var username string
	var address string
	var oscCardNumber uint64
	var totalCost float64
	var itemsEncoded []byte

	err := row.Scan(&uUID, &checkoutDate, &username, &address, &oscCardNumber, &totalCost, &itemsEncoded)
	checkErr(err)

	var items []dao.InventoryItem
	marErr := json.Unmarshal(itemsEncoded, items)
	checkErr(marErr)
	return dao.Purchase{&id, checkoutDate, username, address, oscCardNumber, totalCost, items}
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
