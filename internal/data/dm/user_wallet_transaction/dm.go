package userwallettxndm

import (
	"context"

	"xorm.io/xorm"

	userwallettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	InsertWithSession(
		ctx context.Context,
		session *xorm.Session,
		userWalletTransaction *userwallettxnmodel.UserWalletTransaction,
	) error
}

var (
	defaultDMInstance DM
)

func InitUserWalletTxnDM() {
	defaultDBEngine := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetUserWalletTxnDM() DM {
	return defaultDMInstance
}
