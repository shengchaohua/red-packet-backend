package userdm

import (
	"context"

	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
)

type DM interface {
	Create(ctx context.Context, user *usermodel.User) error
}

func InitUserDM() {
	defaultDM = &DefaultDM{}
}

func GetDataManager() DM {
	return defaultDM
}
