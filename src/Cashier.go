package main

import (
    "Assignment-3/src/dao"
    "Assignment-3/src/dbclient"
)

type Cashier struct {
    user           dao.User
    purchaseClient dbclient.PurchaseDBClient
}
