package dbclient

import (
	"Assignment-3/dao"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UserDBClient struct {
	database *sql.DB
}

var userDBInstance *UserDBClient

func GetUserDBClient() *UserDBClient {
	if userDBInstance == nil {
		userDBInstance = &UserDBClient{initializeDB("users.db")}
		createTable(userDBInstance.database, "users",
			`id Integer primary key,
			checkoutdate Numeric,
			username Text,
			Address Text, 
			oscnum Integer,
			total Real,
			items Blob`)
	}
	return userDBInstance
}

func (client UserDBClient) getUserByUsername(username string) dao.User {
	return dao.User{}
}

func (client UserDBClient) createUser(user dao.User) bool {
	return false
}

func (client UserDBClient) changePassword(user dao.User, password string) bool {
	return false
}

func (client UserDBClient) changeAddress(user dao.User, address string) bool {
	return false
}

func (client UserDBClient) changeOscCardNumber(cardNumber uint64) bool {
	return false
}
