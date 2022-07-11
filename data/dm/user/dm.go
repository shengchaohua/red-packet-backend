package userdm

import (
	"context"

	usermodel "github.com/shengchaohua/red-packet-backend/data/model/user"
)

type DataManager interface {
	Create(ctx context.Context, user *usermodel.User) error
}

func InitDataManger() {
	defaultDM = &DefaultDM{}
}

func GetDataManager() DataManager {
	return defaultDM
}
