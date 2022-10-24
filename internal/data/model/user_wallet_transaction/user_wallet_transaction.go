package redpackettxnmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	UserWalletTransactionTableName = "user_wallet_transaction_tab"
)

type UserWalletTransaction struct {
	*UserWalletTransactionTab
	ExtraData *UserWalletTransactionExtraData
}

type UserWalletTransactionTab struct {
	Id                uint64               `xorm:"'id' bigint pk autoincr"`
	UserId            uint64               `xorm:"'user_id' bigint notnull"`
	TransactionType   enum.TransactionType `xorm:"'transaction_type' int notnull"`
	TransactionStatus uint32               `xorm:"'transaction_status' int notnull"`
	ReferenceId       string               `xorm:"'reference_id' varchar(255) notnull"`
	Amount            uint32               `xorm:"'count' int notnull"`
	Ctime             uint32               `xorm:"'ctime' int notnull"`
	Mtime             uint32               `xorm:"'mtime' int notnull"`
	ExtraData         []byte               `xorm:"'extra_data' blob"`
}

type UserWalletTransactionExtraData struct{}

func (model *UserWalletTransaction) ModelToTab() (*UserWalletTransactionTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user wallet txn model is nil")
	}

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user wallet txn extra data error: %w", err)
	}
	model.UserWalletTransactionTab.ExtraData = extraDataBytes

	return model.UserWalletTransactionTab, nil
}

func (tab *UserWalletTransactionTab) TabToModel() (*UserWalletTransaction, error) {
	if tab == nil {
		return nil, fmt.Errorf("user wallet txn tab is nil")
	}

	model := &UserWalletTransaction{
		UserWalletTransactionTab: tab,
	}

	extraData := &UserWalletTransactionExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user wallet txn extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
