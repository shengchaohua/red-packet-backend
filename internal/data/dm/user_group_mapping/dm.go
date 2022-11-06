package usergroupmappingdm

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergroupmappingmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_mapping"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	LoadByUserIdAndGroupId(
		ctx context.Context,
		userId uint64,
		groupId uint64,
	) (*usergroupmappingmodel.UserGroupMapping, error)
}

var (
	defaultDMInstance DM
)

func InitDM() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	if mainDBEngineManager == nil {
		panic("mainDBEngineManager has not been inited")
	}
	env := config.GetGlobalAppConfig().ServerConfig.GetEnv()
	if env == "" {
		panic("env is empty")
	}
	defaultDMInstance = NewDefaultDM(
		mainDBEngineManager,
		env,
	)
}

func GetDM() DM {
	return defaultDMInstance
}
