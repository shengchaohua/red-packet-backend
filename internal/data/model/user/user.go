package usermodel

// UserTab defines the user table in DB
type UserTab struct {
	UserId   uint64
	Username string
	NickName string
	Ctime    uint32
	Mtime    uint32
}

// User defines the user class
type User struct {
	*UserTab
}
