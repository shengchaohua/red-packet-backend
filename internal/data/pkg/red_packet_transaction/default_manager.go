package redpackettxnpkg

import (
	"context"

	redpackettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	redpackettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet_transaction"
	"xorm.io/xorm"
)

type defaultManager struct {
	redPacketTxnDM redpackettxndm.DM
}

func NewDefaultManager(userWalletTxnDM redpackettxndm.DM) Manager {
	return &defaultManager{
		redPacketTxnDM: userWalletTxnDM,
	}
}

func (manager *defaultManager) AddRedPacketTxn(
	ctx context.Context,
	session *xorm.Session,
	userId uint64,
	transactionType enum.TransactionType,
	referenceId string,
	amount uint32,
) error {
	redPacketTxn := &redpackettxnmodel.RedPacketTransaction{
		RedPacketTransactionTab: &redpackettxnmodel.RedPacketTransactionTab{
			UserId:          userId,
			TransactionType: transactionType,
			ReferenceId:     referenceId,
			Amount:          amount,
		},
	}

	err := manager.redPacketTxnDM.InsertWithSession(ctx, session, redPacketTxn)
	if err != nil {
		return ErrAddUserWalletTxn.WrapWithMsg(err, "insert_new_user_wallet_txn_error")
	}

	return nil
}
