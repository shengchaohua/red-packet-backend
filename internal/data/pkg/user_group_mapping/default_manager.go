package usergroupmappingpkg

import (
	"context"

	usergroupmappingdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_mapping"
)

type defaultManager struct {
	userGroupMappingDM usergroupmappingdm.DM
}

func NewDefaultManager(userGroupMappingDM usergroupmappingdm.DM) Manager {
	return &defaultManager{
		userGroupMappingDM: userGroupMappingDM,
	}
}

func (manager *defaultManager) CheckUserInGroup(
	ctx context.Context,
	userId uint64,
	groupId uint64,
) (bool, error) {
	userGroupMapping, err := manager.userGroupMappingDM.LoadByUserIdAndGroupId(
		ctx,
		userId,
		groupId,
	)
	if err != nil {
		return false, ErrCheckUserInGroup.WrapWithMsg(err, "load_user_group_mapping_error")
	}
	return userGroupMapping != nil, nil
}
