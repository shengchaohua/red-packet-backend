package userservice

import (
	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type Service interface{}

var (
	defaultServiceInstance Service
)

func InitService() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	if mainDBEngineManager == nil {
		panic("mainDBEngineManager has not been inited")
	}
	userManager := userpkg.GetManager()
	if userManager == nil {
		panic("userManager has not been inited")
	}
	userWalletManager := userwalletpkg.GetUserWalletManager()
	if userWalletManager == nil {
		panic("userWalletManager has not been inited")
	}
	userWalletTxnManager := userwallettxnpkg.GetUserWalletTxnManager()
	if userWalletTxnManager == nil {
		panic("userWalletTxnManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(
		mainDBEngineManager,
		userManager,
		userWalletManager,
		userWalletTxnManager,
	)
}

func GetUserService() Service {
	return defaultServiceInstance
}
