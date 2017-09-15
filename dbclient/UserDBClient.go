package dbclient

import "Assignment-3/dao"

type UserDBClient struct {
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

