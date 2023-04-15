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

func (manager *defaultManager) Register(
	ctx context.Context,
	session xorm.Session,
	username string,
	password string,
	email string,
) {
	//TODO implement me
	panic("implement me")
}

func (manager *defaultManager) Login(
	ctx context.Context,
	session *xorm.Session,
	username string,
	password string,
) (*usermodel.User, error) {
	//TODO implement me
	panic("implement me")
}
