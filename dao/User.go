package dao

type User struct {
    Username      string
    Password      string
    Cart          Cart
    Address       string
    OscCardNumber uint64
}
