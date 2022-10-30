package userwallettxndm

import (
	"context"
	"fmt"

	userwallettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/utils"
	"xorm.io/xorm"
)

type defaultDM struct {
	database.EngineManager
	tableName string
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
		tableName:     userwallettxnmodel.UserWalletTransactionTableName}
}

func (dm *defaultDM) InsertWithSession(
	ctx context.Context,
	session *xorm.Session,
	userWalletTransaction *userwallettxnmodel.UserWalletTransaction,
) error {
	if session == nil {
		return ErrParam.WithMsg("session cannot be nil")
	}

	now := utils.GetCurrentTime()
	userWalletTransaction.Ctime = now
	userWalletTransaction.Mtime = now

	userWalletTxnTab, err := userWalletTransaction.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Table(dm.tableName).InsertOne(userWalletTxnTab)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_error|user_id=%d,txn_type=%s,reference_id=%s",
			userWalletTransaction.UserId,
			userWalletTransaction.TransactionType.String(),
			userWalletTransaction.ReferenceId,
		))
	}
	if affected == 0 {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_failed|user_id=%d,txn_type=%s,reference_id=%s",
			userWalletTransaction.UserId,
			userWalletTransaction.TransactionType.String(),
			userWalletTransaction.ReferenceId,
		))
	}

	return nil
}
