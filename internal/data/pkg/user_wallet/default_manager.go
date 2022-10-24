package userwalletpkg

import (
	"context"
	"fmt"

	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
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
	userWallet, err := manager.userWalletDM.LoadByUserIdWithSessionForUpdate(ctx, session, userId)
	if err != nil {
		return ErrDeductUserWallet.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrDeductUserWallet.WithMsg(fmt.Sprintf("user_wallet_not_found|user_id=%d", userId))
	}
	if userWallet.Balance < uint64(amount) {
		return ErrDeductUserWallet
	}

	userWallet.Balance += uint64(amount)
	if err := manager.userWalletDM.UpdateWithSession(ctx, session, userWallet); err != nil {
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
	userWallet, err := manager.userWalletDM.LoadByUserIdWithSessionForUpdate(ctx, session, userId)
	if err != nil {
		return ErrDeductUserWallet.WrapWithMsg(err, "load_user_wallet_error")
	}
	if userWallet == nil {
		return ErrDeductUserWallet.WithMsg(fmt.Sprintf("user_wallet_not_found|user_id=%d", userId))
	}

	userWallet.Balance += uint64(amount)
	if err := manager.userWalletDM.UpdateWithSession(ctx, session, userWallet); err != nil {
		return ErrDeductUserWallet.WrapWithMsg(err, "update_user_wallet_error")
	}

	return nil
}
