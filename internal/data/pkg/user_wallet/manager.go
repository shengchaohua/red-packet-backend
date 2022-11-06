package userwalletpkg

import (
	"context"

	"xorm.io/xorm"

	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
)

type Manager interface {
	CreateUserWallet(
		ctx context.Context,
		session *xorm.Session,
		userId uint64,
	) error

	DeductUserWalletBalance(
		ctx context.Context,
		session *xorm.Session,
		userId uint64,
		amount uint32,
	) error

	AddUserWalletBalance(
		ctx context.Context,
		session *xorm.Session,
		userWalletId uint64,
		amount uint32,
	) error
}

var defaultManagerInstance Manager

func InitManager() {
	userWalletDM := userwalletdm.GetDM()
	defaultManagerInstance = NewDefaultManager(userWalletDM)
}

func GetUserWalletManager() Manager {
	return defaultManagerInstance
}
