package redpacketservice

import (
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultService struct {
	database.EngineManager
	redPacketManager redpacketpkg.Manager
}

func NewDefaultService(
	engineManager database.EngineManager,
	redPacketManager redpacketpkg.Manager,
) Service {
	return &defaultService{
		EngineManager:    engineManager,
		redPacketManager: redPacketManager,
	}
}
