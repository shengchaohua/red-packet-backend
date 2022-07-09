package usermodel

// UserTab defines the user table in DB table
type UserTab struct {
	UserId   uint64
	UserName string
	NickName string
}

// User defines the user class
type User struct {
	*UserTab
}
