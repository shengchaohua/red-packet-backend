package infra

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/infra/database"
)

func InitInfra(ctx context.Context) {
	database.InitDBEngineManager(ctx)
}
