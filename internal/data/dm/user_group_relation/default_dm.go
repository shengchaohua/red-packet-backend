package usergrouprelationdm

import (
	"context"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergrouprelationmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_relation"
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
		tableName:           usergrouprelationmodel.UserGroupRelationTableName,
		shardingNum:         getShardingNumberByEnv(env),
		shardingTableFormat: usergrouprelationmodel.UserGroupRelationShardingTableFormat,
	}
}

func (dm *defaultDM) getShardingTable(userId uint64) string {
	return fmt.Sprintf(dm.shardingTableFormat, userId%dm.shardingNum)
}

func (dm *defaultDM) LoadByUserIdAndGroupId(
	ctx context.Context,
	userId uint64,
	groupId uint64,
) (*usergrouprelationmodel.UserGroupRelation, error) {
	var (
		userGroupMappingTab = &usergrouprelationmodel.UserGroupRelationTab{}
		shardingTable       = dm.getShardingTable(userId)
	)

	has, err := dm.GetSlaveEngine().Table(shardingTable).
		Where("user_id = ? and group_id = ?", userId, groupId).
		Get(userGroupMappingTab)
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
