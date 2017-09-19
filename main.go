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

	addListItemsToShell(storeShell)
	addAddItemToCartToShell(storeShell)

	loginShell.Run()

	if userManagerInstance.user != nil {
		// user managed to log in successfully. Start the store shell.
		storeShell.Run()
	}

	// for loginShell.Active() || storeShell.Active() {

	// }
}
