package usergrouprelationmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	UserRelationTableName           = "user_relation_tab"
	UserRelationShardingTableFormat = UserRelationTableName + "_%08d" // sharded by user_id
)

type UserRelation struct {
	*UserRelationTab
	ExtraData *UserRelationExtraData
}

type UserRelationTab struct {
	Id           uint64                `xorm:"'id' bigint unsigned pk autoincr"`
	UserId       uint64                `xorm:"'user_id' bigint unsigned notnull"`
	FriendId     uint64                `xorm:"'friend_id' bigint unsigned notnull"`
	RelationType enum.UserRelationType `xorm:"'relation_type' int unsigned notnull"`
	Ctime        uint32                `xorm:"'ctime' int unsigned notnull"`
	Mtime        uint32                `xorm:"'mtime' int unsigned notnull"`
	ExtraData    []byte                `xorm:"'extra_data' blob"`
}

type UserRelationExtraData struct{}

func (model *UserRelation) ModelToTab() (*UserRelationTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user relation model is nil")
	}

	tab := model.UserRelationTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user relation extra data error: %w", err)
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

func (tab *UserRelationTab) TabToModel() (*UserRelation, error) {
	if tab == nil {
		return nil, fmt.Errorf("user relation tab is nil")
	}

	model := &UserRelation{
		UserRelationTab: tab,
	}

	extraData := &UserRelationExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user relation extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}
