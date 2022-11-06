package usergrouprelationmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	UserGroupRelationTableName           = "user_group_relation_tab"
	UserGroupRelationShardingTableFormat = UserGroupRelationTableName + "_%08d" // sharded by user_id or group_id
)

type UserGroupRelation struct {
	*UserGroupRelationTab
	ExtraData *UserGroupRelationExtraData
}

type UserGroupRelationTab struct {
	Id           uint64                     `xorm:"'id' bigint unsigned pk autoincr"`
	UserId       uint64                     `xorm:"'user_id' bigint unsigned notnull"`
	GroupId      uint64                     `xorm:"'group_id' bigint unsigned notnull"`
	RelationType enum.UserGroupRelationType `xorm:"'relation_type' int unsigned notnull"`
	Ctime        uint32                     `xorm:"'ctime' int unsigned notnull"`
	Mtime        uint32                     `xorm:"'mtime' int unsigned notnull"`
	ExtraData    []byte                     `xorm:"'extra_data' blob"`
}

type UserGroupRelationExtraData struct{}

func (model *UserGroupRelation) ModelToTab() (*UserGroupRelationTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user group mapping model is nil")
	}

	tab := model.UserGroupRelationTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user group mapping extra data error: %w", err)
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

func (tab *UserGroupRelationTab) TabToModel() (*UserGroupRelation, error) {
	if tab == nil {
		return nil, fmt.Errorf("user group mapping tab is nil")
	}

	model := &UserGroupRelation{
		UserGroupRelationTab: tab,
	}

	extraData := &UserGroupRelationExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user group mapping extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
