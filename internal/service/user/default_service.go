package userservice

import (
	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultService struct {
	database.EngineManager
	userManager      userpkg.Manager
	useWalletManager userwalletpkg.Manager
}

func NewDefaultService(
	engineManager database.EngineManager,
	userManager userpkg.Manager,
	useWalletManager userwalletpkg.Manager,
) *defaultService {
	return &defaultService{
		EngineManager:    engineManager,
		userManager:      userManager,
		useWalletManager: useWalletManager,
	}
}
