package dbclient

import (
	"database/sql"
)

func initializeDB(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	checkErr(err)
	if db == nil {
		panic("DB is nil!")
	}

	return db
}
