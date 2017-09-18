package dao

type Cart struct {
	Items []CartItem
}

type CartItem struct {
	Item     InventoryItem
	Quantity uint64
}

func (cart Cart) AddItem(item InventoryItem, quantity uint64) {
	for i, cartItem := range cart.Items {
		if cartItem.Item == item {
			cart.Items[i].Quantity += quantity
			return
		}
	}

	cart.Items = append(cart.Items, CartItem{item, quantity})
	return
}

func (cart Cart) RemoveItem(item InventoryItem, quantity uint64) {
	var foundItem *CartItem
	var foundIndex int
	for i, cartItem := range cart.Items {
		if cartItem.Item == item {
			foundItem = &cartItem
			foundIndex = i
		}
	}

	if foundItem != nil && foundItem.Quantity-quantity > 0 {
		foundItem.Quantity = foundItem.Quantity - quantity
	} else {
		cart.Items = append(cart.Items[:foundIndex], cart.Items[foundIndex+1:]...)
	}
}

func (cart Cart) GetItems() []InventoryItem {
	var totalQuantity uint64
	for _, item := range cart.Items {
		totalQuantity += item.Quantity
	}

	items := make([]InventoryItem, totalQuantity, totalQuantity)

	for i, item := range cart.Items {
		for j := 0; j < int(item.Quantity); j++ {
			items[i+j] = item.Item
		}
	}

	return items
}

func (cart Cart) EmptyCart() {
	cart.Items = []CartItem{}
}

func (cart Cart) Checkout() {
	// might need a redesign here
}

func (cart Cart) GetTotalCost() float64 {
	var totalCost float64 = 0.0
	for _, item := range cart.Items {
		totalCost += item.Item.Price * float64(item.Quantity)
	}
	return totalCost
}
