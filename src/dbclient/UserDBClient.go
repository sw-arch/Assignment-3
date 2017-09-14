package dbclient

type UserDBClient interface {
	initializeDB() bool
}
