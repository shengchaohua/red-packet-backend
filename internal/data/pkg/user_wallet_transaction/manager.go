package userwallettxnpkg

import (
	"context"

	"xorm.io/xorm"

	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

type Manager interface {
	AddUserWalletTxn(
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
	userWalletTxnDM := userwallettxndm.GetUserWalletTxnDM()
	defaultManagerInstance = NewDefaultManager(userWalletTxnDM)
}

func GetUserWalletTxnManager() Manager {
	return defaultManagerInstance
}
