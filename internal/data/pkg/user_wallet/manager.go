package userwalletpkg

import (
	"context"

	"xorm.io/xorm"

	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
)

type Manager interface {
	DeductUserWallet(
		ctx context.Context,
		session *xorm.Session,
		userId uint64,
		amount uint32,
	) error

	AddUserWallet(
		ctx context.Context,
		session *xorm.Session,
		userWalletId uint64,
		amount uint32,
	) error
}

var defaultManagerInstance Manager

func InitUserWalletManager() {
	userWalletDM := userwalletdm.GetUserWalletDM()
	defaultManagerInstance = NewDefaultManager(userWalletDM)
}

func GetDefaultManager() Manager {
	return defaultManagerInstance
}
