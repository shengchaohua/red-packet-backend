package userwalletdm

import (
	"context"

	"xorm.io/xorm"

	userwalletmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	InsertWithSession(
		ctx context.Context,
		session *xorm.Session,
		userWallet *userwalletmodel.UserWallet,
	) error

	UpdateWithSession(
		ctx context.Context,
		session *xorm.Session,
		userWallet *userwalletmodel.UserWallet,
	) error

	LoadByUserIdWithSessionForUpdate(
		ctx context.Context,
		session *xorm.Session,
		userId uint64,
	) (*userwalletmodel.UserWallet, error)

	LoadByUserId(
		ctx context.Context,
		userId uint64,
		querySlave bool,
		queryMaster bool,
	) (*userwalletmodel.UserWallet, error)
}

var (
	defaultDMInstance DM
)

func InitUserWalletDM() {
	defaultDBEngine := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetUserWalletDM() DM {
	return defaultDMInstance
}
