package userwalletpkg

import (
	"context"

	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
	userwalletmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet"
	"xorm.io/xorm"
)

type defaultManager struct {
	userWalletDM userwalletdm.DM
}

func NewDefaultManager(userWalletDM userwalletdm.DM) Manager {
	return &defaultManager{
		userWalletDM: userWalletDM,
	}
}

func (manager *defaultManager) DeductUserWallet(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
	amount uint32,
) error {
	var (
		userWallet *userwalletmodel.UserWallet
		err        error
	)

	userWallet, err = manager.userWalletDM.LoadByUserIdWithSessionForUpdate(ctx, session, userId)
	if err != nil {
		return ErrDeductUserWallet.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrDeductUserWallet.WithMsg("user_wallet_not_found")
	}
	if userWallet.Balance < uint64(amount) {
		return ErrDeductUserWallet.WithMsg("user_wallet_balance_is_not_enough")
	}

	userWallet.Balance -= uint64(amount)
	err = manager.userWalletDM.UpdateWithSession(ctx, session, userWallet)
	if err != nil {
		return ErrDeductUserWallet.WrapWithMsg(err, "update_user_wallet_error")
	}

	return nil
}

func (manager *defaultManager) AddUserWallet(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
	amount uint32,
) error {
	var (
		userWallet *userwalletmodel.UserWallet
		err        error
	)

	userWallet, err = manager.userWalletDM.LoadByUserIdWithSessionForUpdate(ctx, session, userId)
	if err != nil {
		return ErrAddUserWallet.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrAddUserWallet.WithMsg("user_wallet_not_found")
	}

	userWallet.Balance += uint64(amount)
	err = manager.userWalletDM.UpdateWithSession(ctx, session, userWallet)
	if err != nil {
		return ErrAddUserWallet.WrapWithMsg(err, "update_user_wallet_error")
	}

	return nil
}
