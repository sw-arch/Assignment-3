package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
	"math/rand"
)

type UserManager struct {
	user *dao.User
}

var userManagerInstance *UserManager

func GetUserManager() *UserManager {
	if userManagerInstance == nil {
		userManagerInstance = &UserManager{
			nil}
	}
	return userManagerInstance
}

func (manager *UserManager) logIn(username string, password string) bool {
	user, success := dbclient.GetUserDBClient().GetUserByUsername(username)
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
	if _, userExists := dbclient.GetUserDBClient().GetUserByUsername(username); userExists {
		// Username is taken
		return false
	}

	cardNumber := uint64(rand.Intn(9999999999))

	for _, userExists := dbclient.GetUserDBClient().GetUserByOSCNumber(cardNumber); userExists; _, userExists = dbclient.GetUserDBClient().GetUserByOSCNumber(cardNumber) {
		cardNumber = uint64(rand.Intn(9999999999))
	}

	user := dao.User{
		username,
		password,
		&dao.Cart{},
		address,
		cardNumber}
	created := dbclient.GetUserDBClient().CreateUser(&user)

	if created {
		manager.user = &user
	}

	return created
}

func (manager UserManager) getHistory() []dao.Purchase {
	return dbclient.GetPurchaseDBClient().GetPurchasesByUsername(manager.user.Username)
}
