package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbClient"
)

type Cashier struct {
	user           dao.User
	purchaseClient dbClient.PurchaseDBClient
}

func (cashier Cashier) confirmPurchase() bool {
	return false
}

func (cashier Cashier) previewPurchase() dao.Purchase {
	return dao.Purchase{}
}