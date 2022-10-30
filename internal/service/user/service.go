package userservice

import (
	"context"

	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type Service interface {
	CreateRandomUsers(
		ctx context.Context,
		request *CreateRandomUsersRequest,
	) (*CreateRandomUsersResponse, error)
}

var (
	defaultServiceInstance Service
)

func InitService() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	if mainDBEngineManager == nil {
		panic("mainDBEngineManager has not been inited")
	}
	userManager := userpkg.GetUserManager()
	if userManager == nil {
		panic("userManager has not been inited")
	}
	userWalletManager := userwalletpkg.GetUserWalletManager()
	if userWalletManager == nil {
		panic("userWalletManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(
		mainDBEngineManager,
		userManager,
		userWalletManager,
	)
}

func GetUserService() Service {
	return defaultServiceInstance
}
