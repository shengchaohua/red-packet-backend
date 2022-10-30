package redpacketservice

import (
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultService struct {
	database.EngineManager
	redPacketManager    redpacketpkg.Manager
	userWalletManager   userwalletpkg.Manager
	redPacketTxnManager userwallettxnpkg.Manager
}

func NewDefaultService(
	engineManager database.EngineManager,
	redPacketManager redpacketpkg.Manager,
	userWalletManager userwalletpkg.Manager,
	redPacketTxnManager userwallettxnpkg.Manager,
) Service {
	return &defaultService{
		EngineManager:       engineManager,
		redPacketManager:    redPacketManager,
		userWalletManager:   userWalletManager,
		redPacketTxnManager: redPacketTxnManager,
	}
}
