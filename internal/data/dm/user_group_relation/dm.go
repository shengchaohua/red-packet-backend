package usergrouprelationdm

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergrouprelationmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_relation"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	LoadByUserIdAndGroupId(
		ctx context.Context,
		userId uint64,
		groupId uint64,
	) (*usergrouprelationmodel.UserGroupRelation, error)
}

var (
	defaultDMInstance DM
)

func InitDM() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	env := config.GetGlobalConfig().ServerConfig.EnvEnum
	defaultDMInstance = NewDefaultDM(mainDBEngineManager, env)
}

func GetDM() DM {
	return defaultDMInstance
}
