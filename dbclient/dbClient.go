package dbclient

import (
    "fmt"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

func initializeDB(dbName string) *sql.DB {
    db, err := sql.Open("sqlite3", dbName)
    checkErr(err)
    if db == nil {
        panic("DB is nil!")
    }

    return db
}

func createTable(db *sql.DB, tableName string, schema string) {
    checkStatement := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
    result, err := db.Exec(checkStatement)
    checkErr(err)
    if result == nil {
        createStatement := fmt.Sprintf("Create Table %s (%s);", tableName, schema)
        _, err := db.Exec(createStatement)
        checkErr(err)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
