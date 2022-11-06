package usergrouprelationpkg

import (
	"context"

	usergrouprelationdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_relation"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
)

type defaultManager struct {
	userGroupMappingDM usergrouprelationdm.DM
}

func NewDefaultManager(userGroupMappingDM usergrouprelationdm.DM) Manager {
	return &defaultManager{
		userGroupMappingDM: userGroupMappingDM,
	}
}

func (manager *defaultManager) CheckUserInGroup(
	ctx context.Context,
	userId uint64,
	groupId uint64,
) (bool, error) {
	userGroupRelation, err := manager.userGroupMappingDM.LoadByUserIdAndGroupId(
		ctx,
		userId,
		groupId,
	)
	if err != nil {
		return false, ErrCheckUserInGroup.WrapWithMsg(err, "load_user_group_mapping_error")
	}

	if userGroupRelation != nil &&
		userGroupRelation.RelationType == enum.UserGroupRelationTypeInGroup {
		return true, nil
	}

	return false, nil
}
