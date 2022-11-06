package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	usergroupmappingpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_group_mapping"
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
	redPacketManager := redpacketpkg.GetManager()
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
	userGroupMappingManager := usergroupmappingpkg.GetManager()
	if userGroupMappingManager == nil {
		panic("userGroupMappingManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(
		engineManager,
		redPacketManager,
		userWalletManager,
		userWalletTxnManager,
		userGroupMappingManager,
	)
}

func GetRedPacketService() Service {
	return defaultServiceInstance
}
