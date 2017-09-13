package main

import (
    "Assignment-3/dao"
    "Assignment-3/dbclient"
)

type Cashier struct {
    user           dao.User
    purchaseClient dbclient.PurchaseDBClient
}
