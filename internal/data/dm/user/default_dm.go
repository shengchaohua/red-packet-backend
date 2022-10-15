package userdm

import (
	"context"

	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
)

type DefaultDM struct {
}

var defaultDM *DefaultDM

func (dm *DefaultDM) Create(
	ctx context.Context,
	user *usermodel.User,
) error {
	return nil
}
