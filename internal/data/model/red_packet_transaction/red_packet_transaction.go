package redpackettxnmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	RedPacketTransactionTableName = "red_packet_transaction_tab"
)

type RedPacketTransaction struct {
	*RedPacketTransactionTab
	ExtraData *RedPacketTransactionExtraData
}

type RedPacketTransactionTab struct {
	Id              uint64               `xorm:"'id' bigint pk autoincr"`
	UserId          uint64               `xorm:"'user_id' bigint notnull"`
	TransactionType enum.TransactionType `xorm:"'transaction_type' int notnull"`
	ReferenceId     string               `xorm:"'reference_id' varchar(255) notnull"`
	Amount          uint32               `xorm:"'amount' int notnull"`
	Ctime           uint32               `xorm:"'ctime' int notnull"`
	Mtime           uint32               `xorm:"'mtime' int notnull"`
	ExtraData       []byte               `xorm:"'extra_data' blob"`
}

type RedPacketTransactionExtraData struct{}

func (model *RedPacketTransaction) ModelToTab() (*RedPacketTransactionTab, error) {
	if model == nil {
		return nil, fmt.Errorf("red packet txn model is nil")
	}

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal red packet txn extra data error: %w", err)
	}
	model.RedPacketTransactionTab.ExtraData = extraDataBytes

	return model.RedPacketTransactionTab, nil
}

func (tab *RedPacketTransactionTab) TabToModel() (*RedPacketTransaction, error) {
	if tab == nil {
		return nil, fmt.Errorf("red packet txn tab is nil")
	}

	model := &RedPacketTransaction{
		RedPacketTransactionTab: tab,
	}

	extraData := &RedPacketTransactionExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal red packet txn extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
