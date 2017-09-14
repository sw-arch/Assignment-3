package main

import (
    "Assignment-3/dao"
    "Assignment-3/dbclient"
)

type UserManager struct {
    userClient dbclient.UserDBClient
}

func (manager UserManager) logIn(user dao.User, password string) bool {
    return false
}

func (manager UserManager) logOut(user dao.User) bool {
    return false
}

func (manager UserManager) getHistory(user dao.User) bool {
    return false
}
