package redpackettxnpkg

import (
	"context"

	"xorm.io/xorm"

	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
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

func InitManager() {
	userWalletTxnDM := userwallettxndm.GetDM()
	defaultManagerInstance = NewDefaultManager(userWalletTxnDM)
}

func GetUserWalletTxnManager() Manager {
	if defaultManagerInstance == nil {
		panic("defaultManagerInstance has not been inited")
	}
	return defaultManagerInstance
}
