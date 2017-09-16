package dbclient

import (
    "fmt"
    "Assignment-3/dao"
    "database/sql"

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
            `oscnum Integer primary key,
            username Text,
            password Text,
            address Text,
            cart Blob`)
    }
    return userDBInstance
}

func (client UserDBClient) GetUserByUsername(username string) dao.User {
    return dao.User{}
}

func (client UserDBClient) CreateUser(user dao.User) bool {
    _, err := client.db.Exec(fmt.Sprintf("INSERT INTO users (oscnum, username, password, address, cart) VALUES (%s, %s, %s, %s, %s)", user.OscCardNumber, user.Username, user.Password, user.Address, user.Cart))
    checkErr(err)
    return true
}

func (client UserDBClient) ChangePassword(user dao.User, password string) bool {
    user.Password = password
    _, err := client.db.Exec(fmt.Sprintf("UPDATE users SET password = %s WHERE username == %s", password, user.Username))
    checkErr(err)
    return true
}

func (client UserDBClient) ChangeAddress(user dao.User, address string) bool {
    user.Address = address
    _, err := client.db.Exec(fmt.Sprintf("UPDATE users SET address = %s WHERE username == %s", address, user.Username))
    checkErr(err)
    return true
}

func (client UserDBClient) ChangeOscCardNumber(user dao.User, cardNumber uint64) bool {
    user.OscCardNumber = cardNumber
    _, err := client.db.Exec(fmt.Sprintf("UPDATE users SET oscnum = %s WHERE username == %s", cardNumber, user.Username))
    checkErr(err)
    return true
}
