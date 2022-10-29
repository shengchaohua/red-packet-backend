package userwallettxndm

import (
	"context"
	"fmt"

	redpackettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"xorm.io/xorm"
)

type defaultDM struct {
	database.EngineManager
	tableName string
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
		tableName:     redpackettxnmodel.RedPacketTransactionTableName,
	}
}

func (dm *defaultDM) InsertWithSession(
	ctx context.Context,
	session *xorm.Session,
	redPacketTransaction *redpackettxnmodel.RedPacketTransaction,
) error {
	if session == nil {
		return ErrParam.WithMsg("session cannot be nil")
	}

	redPacketTxnTab, err := redPacketTransaction.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Table(dm.tableName).InsertOne(redPacketTxnTab)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_error|user_id=%d,txn_type=%s,reference_id=%s",
			redPacketTransaction.UserId,
			redPacketTransaction.TransactionType.String(),
			redPacketTransaction.ReferenceId,
		))
	}
	if affected == 0 {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf("insert_db_failed|user_id=%d,txn_type=%s,reference_id=%s",
			redPacketTransaction.UserId,
			redPacketTransaction.TransactionType.String(),
			redPacketTransaction.ReferenceId,
		))
	}

	return nil
}
