package userpkg

import (
	"context"

	"xorm.io/xorm"

	userdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user"
	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
)

type defaultManager struct {
	userDM userdm.DM
}

func NewDefaultManager(userDM userdm.DM) Manager {
	return &defaultManager{
		userDM: userDM,
	}
}

func (manager *defaultManager) CreateUser(
	ctx context.Context,
	session *xorm.Session,
	username string,
) (*usermodel.User, error) {
	user := &usermodel.User{
		UserTab: &usermodel.UserTab{
			Username: username,
		},
	}

	if err := manager.userDM.InsertWithSession(ctx, session, user); err != nil {
		return nil, ErrCreateUser.WrapWithMsg(err, "[CreateUser]create_new_user_error")
	}

	return user, nil
}
