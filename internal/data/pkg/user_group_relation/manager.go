package usergrouprelationpkg

import (
	"context"

	usergrouprelationdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_relation"
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
	userGroupRelationDM := usergrouprelationdm.GetDM()
	if userGroupRelationDM == nil {
		panic("userGroupRelationDM has not been inited")
	}
	defaultManagerInstance = NewDefaultManager(userGroupRelationDM)
}

func GetManager() Manager {
	return defaultManagerInstance
}
