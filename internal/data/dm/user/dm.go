package userdm

import (
	"context"

	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"xorm.io/xorm"
)

type DM interface {
	InsertWithSession(
		ctx context.Context,
		session *xorm.Session,
		user *usermodel.User,
	) error
}

var (
	defaultDMInstance DM
)

func InitDM() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(
		mainDBEngineManager,
	)
}

func GetDM() DM {
	return defaultDMInstance
}
