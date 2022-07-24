package userwalletmodel

// UserTab defines the user table in DB
type UserWalletTab struct {
	UserWalletId uint64
	UserId       uint64
	Balance      uint64
	Ctime        uint64
	Mtime        uint64
}

// User defines the user class
type UserWallet struct {
	*UserWalletTab
}
