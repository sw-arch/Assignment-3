package dbclient

import (
	"Assignment-3/dao"
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type UserDBClient struct {
	db *sql.DB
}

var userDBInstance *UserDBClient

func GetUserDBClient() *UserDBClient {
	if userDBInstance == nil {
		userDBInstance = &UserDBClient{initializeDB("users.db")}
		createTable(userDBInstance.db, "users",
			`username Text PRIMARY KEY,
            password Text,
            cart Blob,
            address Text,
            oscnum Integer UNIQUE`)
	}
	return userDBInstance
}

func (client UserDBClient) GetUserByUsername(username string) (dao.User, bool) {
	statement, prepErr := client.db.Prepare("SELECT * FROM users WHERE username=?")
	checkErr(prepErr)

	var user dao.User
	var cartEncoded []byte
	err := statement.QueryRow(username).Scan(&user.Username, &user.Password, &cartEncoded, &user.Address, &user.OscCardNumber)
	if err == sql.ErrNoRows {
		return dao.User{}, false
	}

	var cart dao.Cart
	marErr := json.Unmarshal(cartEncoded, &cart)
	checkErr(marErr)

	user.PersonalCart = &cart

	return user, true
}

func (client UserDBClient) GetUserByOSCNumber(oscnum uint64) (dao.User, bool) {
	statement, prepErr := client.db.Prepare("SELECT * FROM users WHERE oscnum=?")
	checkErr(prepErr)

	var user dao.User
	var cartEncoded []byte
	err := statement.QueryRow(oscnum).Scan(&user.Username, &user.Password, &cartEncoded, &user.Address, &user.OscCardNumber)
	if err == sql.ErrNoRows {
		return dao.User{}, false
	}

	var cart dao.Cart
	marErr := json.Unmarshal(cartEncoded, &cart)
	checkErr(marErr)

	user.PersonalCart = &cart

	return user, true
}

func (client UserDBClient) CreateUser(user *dao.User) bool {
	encodedCart, marshalErr := json.Marshal(user.PersonalCart)
	checkErr(marshalErr)
	statement, prepErr := client.db.Prepare("INSERT INTO users (username, password, cart, address, oscnum) VALUES (?, ?, ?, ?, ?);")
	checkErr(prepErr)
	_, err := statement.Exec(user.Username, user.Password, encodedCart, user.Address, user.OscCardNumber)
	checkErr(err)
	return true
}

func (client UserDBClient) SaveCart(user *dao.User) bool {
	encodedCart, marshalErr := json.Marshal(user.PersonalCart)
	checkErr(marshalErr)
	statement, prepErr := client.db.Prepare("UPDATE users SET cart = ? WHERE username == ?;")
	checkErr(prepErr)
	_, err := statement.Exec(encodedCart, user.Username)
	checkErr(err)
	return true
}

func (client UserDBClient) ChangePassword(user *dao.User, password string) bool {
	user.Password = password
	statement, prepErr := client.db.Prepare("UPDATE users SET password = ? WHERE username == ?;")
	checkErr(prepErr)
	_, err := statement.Exec(password, user.Username)
	checkErr(err)
	return true
}

func (client UserDBClient) ChangeAddress(user *dao.User, address string) bool {
	user.Address = address
	statement, prepErr := client.db.Prepare("UPDATE users SET address = ? WHERE username == ?;")
	checkErr(prepErr)
	_, err := statement.Exec(address, user.Username)
	checkErr(err)
	return true
}

func (client UserDBClient) ChangeOscCardNumber(user *dao.User, cardNumber uint64) bool {
	user.OscCardNumber = cardNumber
	statement, prepErr := client.db.Prepare("UPDATE users SET oscnum = ? WHERE username == ?;")
	checkErr(prepErr)
	_, err := statement.Exec(cardNumber, user.Username)
	checkErr(err)
	return true
}
