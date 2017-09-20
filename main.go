package main

import (
	"Assignment-3/manager"
	"Assignment-3/ui"

	"github.com/abiosoft/ishell"
)

func main() {
	loginShell := ishell.New()
	storeShell := ishell.New()

	loginShell.Println("Welcome to the Store")

	ui.AddLoginToShell(loginShell)
	ui.AddRegisterToShell(loginShell)
	ui.AddListItemsToShell(loginShell)

	ui.AddListItemsToShell(storeShell)
	ui.AddAddItemToCartToShell(storeShell)
	ui.AddRemoveItemFromCartToShell(storeShell)
	ui.AddDisplayCartToShell(storeShell)
	ui.AddCheckoutToShell(storeShell)
	ui.AddPurchaseHistoryToShell(storeShell)
	ui.AddWhoamiToShell(storeShell)
	ui.AddLogoutToShell(storeShell)

login:
	loginShell.Run()

	if manager.GetUserManager().User != nil {
		// user managed to log in successfully. Start the store shell.
		storeShell.Run()
		if manager.GetUserManager().User == nil {
			// user logged out, restart the login shell.
			goto login
		}
	}
}
