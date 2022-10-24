package userwalletmodel

import (
	"encoding/json"
	"fmt"
)

const (
	UserWalletTableName = "user_wallet_tab"
)

// User defines the user class
type UserWallet struct {
	*UserWalletTab
	ExtraData *UserWalletExtraData
}

// UserTab defines the user table in DB
type UserWalletTab struct {
	Id        uint64 `xorm:"'id' bigint pk autoincr"`
	UserId    uint64 `xorm:"'user_id' bigint notnull"`
	Balance   uint64 `xorm:"'balance' bigint notnull"`
	Ctime     uint32 `xorm:"'ctime' int notnull"`
	Mtime     uint32 `xorm:"'mtime' int notnull"`
	ExtraData []byte `xorm:"'extra_data' blob"`
}

type UserWalletExtraData struct{}

// ModelToTab converts model to tab
func (model *UserWallet) ModelToTab() (*UserWalletTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user wallet model is nil")
	}

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user wallet extra data error: %w", err)
	}
	model.UserWalletTab.ExtraData = extraDataBytes

	return model.UserWalletTab, nil
}

// TabToModel converts tab to model
func (tab *UserWalletTab) TabToModel() (*UserWallet, error) {
	if tab == nil {
		return nil, fmt.Errorf("user wallet tab is nil")
	}

	model := &UserWallet{
		UserWalletTab: tab,
	}

	extraData := &UserWalletExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user wallet extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
