package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
)

type UserManager struct {
	userClient dbclient.UserDBClient
    purchaseClient dbclient.PurchaseDBClient
}

func (manager UserManager) logIn(username string, password string) (dao.User, bool) {
    user := manager.userClient.GetUserByUsername(username)
	if user.Password == password {
        return user, true
    } else {
        return dao.User{}, false
    }
}

func (manager UserManager) logOut(user dao.User) bool {
	return false
}

func (manager UserManager) getHistory(user dao.User) []dao.Purchase {
	return []dao.Purchase{}
}
