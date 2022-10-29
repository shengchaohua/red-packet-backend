package userservice

import (
	"context"

	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
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
	userWalletManager := userwalletpkg.GetDefaultManager()
	if userWalletManager == nil {
		panic("userWalletManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(userWalletManager)
}

func GetService() Service {
	return defaultServiceInstance
}
