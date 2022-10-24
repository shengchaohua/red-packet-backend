package redpacketservice

import (
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	redpackettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet_transaction"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultService struct {
	database.EngineManager
	redPacketManager    redpacketpkg.Manager
	userWalletManager   userwalletpkg.Manager
	redPacketTxnManager redpackettxnpkg.Manager
}

func NewDefaultService(
	engineManager database.EngineManager,
	redPacketManager redpacketpkg.Manager,
	userWalletManager userwalletpkg.Manager,
	redPacketTxnManager redpackettxnpkg.Manager,
) Service {
	return &defaultService{
		EngineManager:       engineManager,
		redPacketManager:    redPacketManager,
		userWalletManager:   userWalletManager,
		redPacketTxnManager: redPacketTxnManager,
	}
}
