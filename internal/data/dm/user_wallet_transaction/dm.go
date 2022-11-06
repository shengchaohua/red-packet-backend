package userwallettxndm

import (
	"context"

	userwallettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"xorm.io/xorm"
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

func InitDM() {
	defaultDBEngine := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetDM() DM {
	return defaultDMInstance
}
