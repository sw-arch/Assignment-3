package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
)

type Store struct{}

type Category struct {
	Identifier  string
	Description string
}

func GetStore() *Store {
	return &Store{}
}

func (store *Store) GetCategories() []Category {
	categories := []Category{}
	for _, category := range dbclient.GetInventoryDBClient().GetCategoryDescriptions() {
		categories = append(categories, Category{category[0], category[1]})
	}
	return categories
}

func (store *Store) GetItemsInCategory(category string) []dao.InventoryItem {
	return dbclient.GetInventoryDBClient().GetItemsByCategory(category)
}

func (store *Store) GetAttributesByCategory(category string) (string, string) {
	return dbclient.GetInventoryDBClient().GetAttributesByCategory(category)
}

func (store *Store) AddToCart(item dao.InventoryItem, quantityDesired int) {
	dbclient.GetInventoryDBClient().Reserve(item, int64(quantityDesired))
	GetUserManager().user.PersonalCart.AddItem(item, int64(quantityDesired))
	dbclient.GetUserDBClient().SaveCart(GetUserManager().user)
}

func (store *Store) RemoveFromCart(item dao.InventoryItem, quantityToRemove int) {
	dbclient.GetInventoryDBClient().Release(item, int64(quantityToRemove))
	GetUserManager().user.PersonalCart.RemoveItem(item, int64(quantityToRemove))
	dbclient.GetUserDBClient().SaveCart(GetUserManager().user)
}
