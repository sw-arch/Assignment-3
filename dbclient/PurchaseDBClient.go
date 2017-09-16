package dbclient

import (
    "Assignment-3/dao"
    "database/sql"

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
            `id Integer primary key,
            checkoutdate Numeric,
            username Text,
            address Text,
            oscnum Integer,
            total Real,
            items Blob`)
    }
    return purchaseDBInstance
}

func (client PurchaseDBClient) getPurchaseByID(id uuid.UUID) dao.Purchase {
    return dao.Purchase{}
}

func (client PurchaseDBClient) addPurchase(purchase dao.Purchase) bool {
    return false
}
