package manager

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
	"math/rand"
)

type UserManager struct {
	User *dao.User
}

var userManagerInstance *UserManager

func GetUserManager() *UserManager {
	if userManagerInstance == nil {
		userManagerInstance = &UserManager{
			nil}
	}
	return userManagerInstance
}

func (manager *UserManager) LogIn(username string, password string) bool {
	user, success := dbclient.GetUserDBClient().GetUserByUsername(username)
	if success && user.Password == password {
		manager.User = &user
		return true
	}
	return false
}

func (manager *UserManager) LogOut() {
	manager.User = nil
	return
}

func (manager *UserManager) Register(username string, password string, address string) bool {
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
		manager.User = &user
	}

	return created
}

func (manager *UserManager) ChangeAddress(newAddress string) {
	dbclient.GetUserDBClient().ChangeAddress(manager.User, newAddress)
}

func (manager *UserManager) GetHistory() []dao.Purchase {
	return dbclient.GetPurchaseDBClient().GetPurchasesByUsername(manager.User.Username)
}
