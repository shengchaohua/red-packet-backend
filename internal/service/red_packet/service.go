package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type Service interface {
	CreateRedPacket(
		ctx context.Context,
		request *CreateRedPacketRequest,
	) (*CreateRedPacketResponse, error)
}

var defaultServiceInstance Service

func InitRedPacketService() {
	engineManager := database.GetMainDBEngineManager()
	if engineManager == nil {
		panic("engineManager has not been inited")
	}
	redPacketManager := redpacketpkg.GetDefaultManager()
	if redPacketManager == nil {
		panic("redPacketManager has not been inited")
	}
	defaultServiceInstance = NewDefaultService(
		engineManager,
		redPacketManager,
	)
}

func GetRedPacketService() Service {
	return defaultServiceInstance
}
