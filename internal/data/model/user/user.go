package usermodel

import "github.com/shengchaohua/red-packet-backend/internal/data/enum"

// User defines the user class
type User struct {
	*UserTab
}

// UserTab defines the user table in DB
type UserTab struct {
	Id        uint64 `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	UserType  enum.RedPacketCategory
	NickName  string `json:"nick_name,omitempty"`
	Ctime     uint32 `json:"ctime,omitempty"`
	Mtime     uint32 `json:"mtime,omitempty"`
	ExtraData []byte `json:"extra_data,omitempty"`
}
