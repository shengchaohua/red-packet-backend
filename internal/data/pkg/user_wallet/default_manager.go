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

func (manager *defaultManager) CreateUserWallet(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
) error {
	userWallet := &userwalletmodel.UserWallet{
		UserWalletTab: &userwalletmodel.UserWalletTab{
			UserId: userId,
		},
	}

	err := manager.userWalletDM.InsertWithSession(ctx, session, userWallet)
	if err != nil {
		return ErrDeductUserWalletBalance.WrapWithMsg(err, "create_user_wallet_error")
	}

	return nil
}

func (manager *defaultManager) DeductUserWalletBalance(
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
		return ErrDeductUserWalletBalance.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrDeductUserWalletBalance.WithMsg("user_wallet_not_found")
	}
	if userWallet.Balance < uint64(amount) {
		return ErrDeductUserWalletBalance.WithMsg("user_wallet_balance_is_not_enough")
	}

	userWallet.Balance -= uint64(amount)
	err = manager.userWalletDM.UpdateWithSession(ctx, session, userWallet)
	if err != nil {
		return ErrDeductUserWalletBalance.WrapWithMsg(err, "update_user_wallet_error")
	}

	return nil
}

func (manager *defaultManager) AddUserWalletBalance(
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
		return ErrAddUserWalletBalance.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrAddUserWalletBalance.WithMsg("user_wallet_not_found")
	}

	userWallet.Balance += uint64(amount)
	err = manager.userWalletDM.UpdateWithSession(ctx, session, userWallet)
	if err != nil {
		return ErrAddUserWalletBalance.WrapWithMsg(err, "update_user_wallet_error")
	}

	return nil
}
