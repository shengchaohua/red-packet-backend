package usergroupmappingpkg

import (
	"context"

	usergroupmappingdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_mapping"
)

type Manager interface {
	CheckUserInGroup(
		ctx context.Context,
		userId uint64,
		groupId uint64,
	) (bool, error)
}

var (
	defaultManagerInstance Manager
)

func InitManager() {
	userGroupMappingDM := usergroupmappingdm.GetDM()
	defaultManagerInstance = NewDefaultManager(userGroupMappingDM)
}

func GetManager() Manager {
	return defaultManagerInstance
}
