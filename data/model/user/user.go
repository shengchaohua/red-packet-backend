package usermodel

// UserTab defines the user table in DB
type UserTab struct {
	UserId   uint64
	Username string
	NickName string
	Ctime    uint64
	Mtime    uint64
}

// User defines the user class
type User struct {
	*UserTab
}
