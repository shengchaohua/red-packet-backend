package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type Service interface {
	CreateRedPacket(
		ctx context.Context,
		request *CreateRedPacketRequest,
	) (*CreateRedPacketResponse, error)

	OpenRedPacket(
		ctx context.Context,
		request *OpenRedPacketRequest,
	) (*OpenRedPacketRequest, error)
}

var defaultServiceInstance Service

func InitService() {
	engineManager := database.GetMainDBEngineManager()
	if engineManager == nil {
		panic("engineManager has not been inited")
	}
	redPacketManager := redpacketpkg.GetRedPacketManager()
	if redPacketManager == nil {
		panic("redPacketManager has not been inited")
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
		engineManager,
		redPacketManager,
		userWalletManager,
		userWalletTxnManager,
	)
}

func GetRedPacketService() Service {
	return defaultServiceInstance
}
