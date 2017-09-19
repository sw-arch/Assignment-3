package dao

type User struct {
	Username      string
	Password      string
	PersonalCart  *Cart
	Address       string
	OscCardNumber uint64
}
