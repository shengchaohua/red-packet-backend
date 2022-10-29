package redpackettxnpkg

import (
	"context"

	"xorm.io/xorm"

	redpackettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

type Manager interface {
	AddRedPacketTxn(
		ctx context.Context,
		session *xorm.Session,
		userId uint64,
		transactionType enum.TransactionType,
		referenceId string,
		amount uint32,
	) error
}

var defaultManagerInstance Manager

func InitRedPacketTxnManager() {
	redPacketTxnDM := redpackettxndm.GetRedPacketTxnDM()
	defaultManagerInstance = NewDefaultManager(redPacketTxnDM)
}

func GetManager() Manager {
	return defaultManagerInstance
}
