package userwalletmodel

// User defines the user class
type UserWallet struct {
	*UserWalletTab
}

// UserTab defines the user table in DB
type UserWalletTab struct {
	UserWalletId uint64
	UserId       uint64
	Balance      uint64
	Ctime        uint32
	Mtime        uint32
}
