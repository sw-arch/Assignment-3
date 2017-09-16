package main

import "Assignment-3/dao"

type Cart struct {
    items     dao.Cart
    store     Store
}

func (cart Cart) AddItem(item dao.InventoryItem, quantity uint64) bool {
    cart.items[item] = quantity
    return false  // true if added
}

func (cart Cart) RemoveItem(item dao.InventoryItem, quantity uint64) {
    quantityInCart, found := cart.items[item]
    if remainder := quantityInCart - quantity; found && remainder > 0 {
        cart.items[item] = remainder
    } else {
        delete(cart.items, item)
    }
}

func (cart Cart) GetItems() []dao.InventoryItem {
    var totalQuantity uint64
    for _, quantity := range cart.items {
        totalQuantity += quantity
    }

    items := make([]dao.InventoryItem, totalQuantity, totalQuantity)

    var i uint64
    for item, quantity := range cart.items {
        for j := uint64(0); j < quantity; j++ {
            items[i + j] = item
        }
    }

    return items
}

func (cart Cart) EmptyCart() {
    cart.items = make(map[dao.InventoryItem]uint64)
}

func (cart Cart) Checkout() {
    // might need a redesign here
}

func (cart Cart) GetTotalCost() float64 {
    var totalCost float64 = 0.0
    for item, quantity := range cart.items {
        totalCost += item.Price * float64(quantity)
    }
    return totalCost
}
