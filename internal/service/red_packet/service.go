package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	redpackettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet_transaction"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
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

func InitRedPacketService() {
	engineManager := database.GetMainDBEngineManager()
	if engineManager == nil {
		panic("engineManager has not been inited")
	}
	redPacketManager := redpacketpkg.GetRedPacketManager()
	if redPacketManager == nil {
		panic("redPacketManager has not been inited")
	}
	userWalletManager := userwalletpkg.GetDefaultManager()
	if userWalletManager == nil {
		panic("userWalletManager has not been inited")
	}
	redPacketTxnManager := redpackettxnpkg.GetManager()
	if redPacketTxnManager == nil {
		panic("redPacketTxnManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(
		engineManager,
		redPacketManager,
		userWalletManager,
		redPacketTxnManager,
	)
}

func GetService() Service {
	return defaultServiceInstance
}
