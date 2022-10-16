package internalpkg

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

func InitPkg(ctx context.Context) {
	database.InitEngineManager(ctx)
}
