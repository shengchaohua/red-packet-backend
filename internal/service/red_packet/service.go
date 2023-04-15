package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	usergrouprelationpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_group_relation"
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
	redPacketManager := redpacketpkg.GetManager()
	userWalletManager := userwalletpkg.GetUserWalletManager()
	userWalletTxnManager := userwallettxnpkg.GetUserWalletTxnManager()
	userGroupRelationManager := usergrouprelationpkg.GetManager()
	defaultServiceInstance = NewDefaultService(
		engineManager,
		redPacketManager,
		userWalletManager,
		userWalletTxnManager,
		userGroupRelationManager,
	)
}

func GetRedPacketService() Service {
	if defaultServiceInstance == nil {
		panic("defaultServiceInstance is nil")
	}
	return defaultServiceInstance
}
