package userservice

import userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"

type defaultService struct {
	useWalletManager userwalletpkg.Manager
}

func NewDefaultService(useWalletManager userwalletpkg.Manager) *defaultService {
	return &defaultService{
		useWalletManager: useWalletManager,
	}
}
