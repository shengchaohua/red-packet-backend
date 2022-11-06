package userpkg

import (
	"context"

	"xorm.io/xorm"

	userdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user"
	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
)

type Manager interface {
	CreateUser(
		ctx context.Context,
		session *xorm.Session,
		username string,
	) (*usermodel.User, error)
}

var (
	defaultManagerInstance Manager
)

func InitManager() {
	userDM := userdm.GetDM()
	if userDM == nil {
		panic("userDM has not been inited")
	}
	defaultManagerInstance = NewDefaultManager(userDM)
}

func GetManager() Manager {
	return defaultManagerInstance
}
