package dbClient

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
		userDBInstance = &UserDBClient{initializeDB()}
		userDBInstance.createUserTable()
	}
	return userDBInstance
}

func initializeUserDB() *sql.DB {
	db, err := sql.Open("sqlite3", "users.db")
	checkErr(err)
	if db == nil {
		panic("DB is nil!")
	}

	return db
}

func (client UserDBClient) createUserTable() {
	checkStatement := "SELECT name FROM sqlite_master WHERE type='table' AND name='user';"
	result, err := client.database.Exec(checkStatement)
	checkErr(err)
	if result == nil {
		createStatement := "Create Table user (id Integer primary key, checkoutdate Numeric, username Text, Address Text, oscnum Integer, total Real, items Blob);"
		_, err := client.database.Exec(createStatement)
		checkErr(err)
	}
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
