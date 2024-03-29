package usermodel

import (
	"encoding/json"
	"fmt"
)

const (
	UserTableName = "user_tab"
)

type User struct {
	*UserTab
	ExtraData *UserExtraData
}

type UserTab struct {
	Id        uint64 `xorm:"'id' bigint unsigned pk autoincr"`
	Username  string `xorm:"'user_name' varchar(255)"`
	Ctime     uint32 `xorm:"'ctime' int unsigned notnull"`
	Mtime     uint32 `xorm:"'mtime' int unsigned notnull"`
	ExtraData []byte `xorm:"'extra_data' blob"`
}

type UserExtraData struct{}

func (model *User) ModelToTab() (*UserTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user model is nil")
	}

	tab := model.UserTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user extra data error: %w", err)
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

func (tab *UserTab) TabToModel() (*User, error) {
	if tab == nil {
		return nil, fmt.Errorf("user tab is nil")
	}

	model := &User{
		UserTab: tab,
	}

	extraData := &UserExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
