package usergroupmappingdm

import (
	"context"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergroupmappingmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_mapping"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultDM struct {
	database.EngineManager
	tableName           string
	shardingNum         uint64
	shardingTableFormat string
}

func NewDefaultDM(engineManager database.EngineManager, env config.Env) DM {
	return &defaultDM{
		EngineManager:       engineManager,
		tableName:           usergroupmappingmodel.UserGroupMappingTableName,
		shardingNum:         getShardingNumberByEnv(env),
		shardingTableFormat: usergroupmappingmodel.UserGroupMappingShardingTableFormat,
	}
}

func (dm *defaultDM) getShardingTable(userIdOrGroupId uint64) string {
	return fmt.Sprintf(dm.shardingTableFormat, userIdOrGroupId%dm.shardingNum)
}

func (dm *defaultDM) LoadByUserIdAndGroupId(
	ctx context.Context,
	userId uint64,
	groupId uint64,
) (*usergroupmappingmodel.UserGroupMapping, error) {
	shardingTable := dm.getShardingTable(userId)
	userGroupMappingTab := &usergroupmappingmodel.UserGroupMappingTab{}
	has, err := dm.GetSlaveEngine().Table(shardingTable).Get(userGroupMappingTab)
	if err != nil {
		return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf("query_db_error|user_id=%d,group_id=%d",
			userId, groupId,
		))
	}
	if !has {
		return nil, nil
	}

	return userGroupMappingTab.TabToModel()
}
