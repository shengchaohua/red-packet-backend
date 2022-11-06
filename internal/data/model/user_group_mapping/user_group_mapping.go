package usergroupmappingmodel

import (
	"encoding/json"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

const (
	UserGroupMappingTableName           = "user_group_mapping_tab"
	UserGroupMappingShardingTableFormat = UserGroupMappingTableName + "_{%08d}" // sharded by user_id and group_id

	shardingNumberTestEnv = 10
	shardingNumberLiveEnv = 1000
)

// User defines the user class
type UserGroupMapping struct {
	*UserGroupMappingTab
	ExtraData *UserGroupMappingExtraData
}

// UserGroupMappingTab defines the user table in DB
type UserGroupMappingTab struct {
	Id        uint64 `xorm:"'id' bigint unsigned pk autoincr"`
	UserId    uint64 `xorm:"'user_id' bigint unsigned notnull"`
	GroupId   uint64 `xorm:"'group_id' bigint unsigned notnull"`
	Ctime     uint32 `xorm:"'ctime' int unsigned notnull"`
	Mtime     uint32 `xorm:"'mtime' int unsigned notnull"`
	ExtraData []byte `xorm:"'extra_data' blob"`
}

type UserGroupMappingExtraData struct{}

func (model *UserGroupMapping) ModelToTab() (*UserGroupMappingTab, error) {
	if model == nil {
		return nil, fmt.Errorf("user group mapping model is nil")
	}

	tab := model.UserGroupMappingTab

	extraDataBytes, err := json.Marshal(model.ExtraData)
	if err != nil {
		return nil, fmt.Errorf("marshal user group mapping extra data error: %w", err)
	}
	tab.ExtraData = extraDataBytes

	return tab, nil
}

func (tab *UserGroupMappingTab) TabToModel() (*UserGroupMapping, error) {
	if tab == nil {
		return nil, fmt.Errorf("user group mapping tab is nil")
	}

	model := &UserGroupMapping{
		UserGroupMappingTab: tab,
	}

	extraData := &UserGroupMappingExtraData{}
	if err := json.Unmarshal(tab.ExtraData, extraData); err != nil {
		return nil, fmt.Errorf("unmarshal user group mapping extra data error: %w", err)
	}
	model.ExtraData = extraData

	return model, nil
}

func GetShardingTableByUserIdAnfGroupId(userId uint64, group_id uint64) (string, error) {
	var (
		shardingIndex uint64

	)
	if userId != 0 && group_id != 0 {
		if 
	}
	if userId != 0 {
		shardingIndex
	}
}

func getShardingTable(userId uint64, group_id uint64, env config.Env) (string, error) {
		var (
		shardingIndex uint64
		shardingNumber uint64
	)

}
