package userwallettxnpkg

import (
	"context"

	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	userwallettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet_transaction"
	"xorm.io/xorm"
)

type defaultManager struct {
	redPacketTxnDM userwallettxndm.DM
}

func NewDefaultManager(userWalletTxnDM userwallettxndm.DM) Manager {
	return &defaultManager{
		redPacketTxnDM: userWalletTxnDM,
	}
}

func (manager *defaultManager) AddUserWalletTxn(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
	transactionType enum.TransactionType,
	referenceId string,
	amount uint32,
) error {
	userWalletTxn := &userwallettxnmodel.UserWalletTransaction{
		UserWalletTransactionTab: &userwallettxnmodel.UserWalletTransactionTab{
			UserId:          userId,
			TransactionType: transactionType,
			ReferenceId:     referenceId,
			Amount:          amount,
		},
	}

	err := manager.redPacketTxnDM.InsertWithSession(ctx, session, userWalletTxn)
	if err != nil {
		return ErrAddUserWalletTxn.WrapWithMsg(err, "insert_new_user_wallet_txn_error")
	}

	return nil
}
