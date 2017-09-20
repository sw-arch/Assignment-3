package main

import (
	"github.com/abiosoft/ishell"
)

func main() {
	loginShell := ishell.New()
	storeShell := ishell.New()

	loginShell.Println("Welcome to the Store")

	addLoginToShell(loginShell)
	addRegisterToShell(loginShell)
	addListItemsToShell(loginShell)

	addListItemsToShell(storeShell)
	addAddItemToCartToShell(storeShell)
	addRemoveItemFromCartToShell(storeShell)
	addDisplayCartToShell(storeShell)
	addCheckoutToShell(storeShell)
	addPurchaseHistoryToShell(storeShell)
	addWhoamiToShell(storeShell)
	addLogoutToShell(storeShell)

login:
	loginShell.Run()

	if GetUserManager().user != nil {
		// user managed to log in successfully. Start the store shell.
		storeShell.Run()
		if GetUserManager().user == nil {
			// user logged out, restart the login shell.
			goto login
		}
	}
}
