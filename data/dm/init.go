package dm

import (
	"context"

	userdm "github.com/shengchaohua/red-packet-backend/data/dm/user"
)

func InitDM(ctx context.Context) {
	userdm.InitUserDM(ctx)
}
