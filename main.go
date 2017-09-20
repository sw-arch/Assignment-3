package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"

	"github.com/abiosoft/ishell"
)

var userManagerInstance *UserManager

func GetUserManager() *UserManager {
	if userManagerInstance == nil {
		userManagerInstance = &UserManager{
			dbclient.GetUserDBClient(),
			dbclient.GetPurchaseDBClient(),
			&dao.User{}}
	}
	return userManagerInstance
}

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

	if userManagerInstance != nil && userManagerInstance.user != nil {
		// user managed to log in successfully. Start the store shell.
		storeShell.Run()
		if userManagerInstance.user.Username == "" {
			goto login
		}
	}

	// for loginShell.Active() || storeShell.Active() {

	// }
}
