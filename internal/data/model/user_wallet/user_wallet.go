package userwalletmodel

// User defines the user class
type UserWallet struct {
	*UserWalletTab
}

// UserTab defines the user table in DB
type UserWalletTab struct {
	UserWalletId uint64 `json:"user_wallet_id,omitempty"`
	UserId       uint64 `json:"user_id,omitempty"`
	Balance      uint64 `json:"balance,omitempty"`
	Ctime        uint32 `json:"ctime,omitempty"`
	Mtime        uint32 `json:"mtime,omitempty"`
}
