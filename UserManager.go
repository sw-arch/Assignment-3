package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
	"math/rand"
)

type UserManager struct {
	userClient     *dbclient.UserDBClient
	purchaseClient *dbclient.PurchaseDBClient
	user           *dao.User
}

func (manager *UserManager) logIn(username string, password string) bool {
	user, success := manager.userClient.GetUserByUsername(username)
	if success && user.Password == password {
		manager.user = &user
		return true
	}
	return false
}

func (manager *UserManager) logOut() {
	manager.user = nil
	return
}

func (manager UserManager) register(username string, password string, address string) bool {
	if _, userExists := manager.userClient.GetUserByUsername(username); userExists {
		// Username is taken
		return false
	}

	cardNumber := uint64(rand.Intn(9999999999))

	for _, userExists := manager.userClient.GetUserByOSCNumber(cardNumber); userExists; _, userExists = manager.userClient.GetUserByOSCNumber(cardNumber) {
		cardNumber = uint64(rand.Intn(9999999999))
	}

	user := dao.User{
		username,
		password,
		&dao.Cart{},
		address,
		cardNumber}
	created := manager.userClient.CreateUser(&user)

	if created {
		manager.user = &user
	}

	return created
}

func (manager UserManager) getHistory(user dao.User) []dao.Purchase {
	return []dao.Purchase{}
}
