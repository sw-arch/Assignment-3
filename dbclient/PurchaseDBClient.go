package dbclient

import (
	"Assignment-3/dao"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
)

type PurchaseDBClient struct {
	database *sql.DB
}

var purchaseDBInstance *PurchaseDBClient

func GetPurchaseDBClient() *PurchaseDBClient {
	if purchaseDBInstance == nil {
		purchaseDBInstance = &PurchaseDBClient{initializeDB("purchase.db")}
		purchaseDBInstance.createPurchaseTable()
	}
	return purchaseDBInstance
}

func (client PurchaseDBClient) createPurchaseTable() {
	checkStatement := "SELECT name FROM sqlite_master WHERE type='table' AND name='purchase';"
	result, err := client.database.Exec(checkStatement)
	checkErr(err)
	if result == nil {
		createStatement := "Create Table purchase (id Integer primary key, checkoutdate Numeric, username Text, Address Text, oscnum Integer, total Real, items Blob);"
		_, err := client.database.Exec(createStatement)
		checkErr(err)
	}
}

func (client PurchaseDBClient) getPurchaseByID(id uuid.UUID) dao.Purchase {
	return dao.Purchase{}
}

func (client PurchaseDBClient) addPurchase(purchase dao.Purchase) bool {
	return false
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
