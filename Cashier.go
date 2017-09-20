package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
	"time"

	"github.com/satori/go.uuid"
)

type Cashier struct{}

func GetCashier() *Cashier {
	return &Cashier{}
}

func (cashier *Cashier) confirmPurchase(user *dao.User, purchase *dao.Purchase) bool {
	purchase.CheckoutDate = time.Now()
	for _, cartItem := range user.PersonalCart.Items {
		dbclient.GetInventoryDBClient().Remove(cartItem.Item, cartItem.Quantity)
	}
	dbclient.GetPurchaseDBClient().AddPurchase(purchase)
	user.PersonalCart.EmptyCart()
	dbclient.GetUserDBClient().SaveCart(user)
	return true
}

func (cashier Cashier) createPurchase(user *dao.User) dao.Purchase {
	return dao.Purchase{uuid.NewV4(), time.Now(), user.Username, user.Address, user.OscCardNumber, user.PersonalCart.GetTotalCost(), user.PersonalCart}
}
