package userservice

import (
	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultService struct {
	database.EngineManager
	userManager          userpkg.Manager
	userWalletManager    userwalletpkg.Manager
	userWalletTxnManager userwallettxnpkg.Manager
}

func NewDefaultService(
	engineManager database.EngineManager,
	userManager userpkg.Manager,
	useWalletManager userwalletpkg.Manager,
	userWalletTxnManager userwallettxnpkg.Manager,
) *defaultService {
	return &defaultService{
		EngineManager:        engineManager,
		userManager:          userManager,
		userWalletManager:    useWalletManager,
		userWalletTxnManager: userWalletTxnManager,
	}
}
