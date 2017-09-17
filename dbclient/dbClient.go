package dbclient

import (
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
	selectStatement, prepSelectErr := db.Prepare("SELECT name FROM sqlite_master WHERE type='table' AND name='?';")
	checkErr(prepSelectErr)

	result, selectErr := selectStatement.Exec(tableName)
	checkErr(selectErr)
	if result == nil {
		createStatement, prepCreateErr := db.Prepare("Create Table ? (?);")
		checkErr(prepCreateErr)
		_, createErr := createStatement.Exec(tableName, schema)
		checkErr(createErr)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
