package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
)

type Service interface {
	CreateRedPacket(
		ctx context.Context,
		request *CreateRedPacketRequest,
	) (*CreateRedPacketResponse, error)
}

var defaultServiceInstance Service

func InitRedPacketService() {
	redPacketManager := redpacketpkg.GetDefaultManager()
	defaultServiceInstance = NewDefaultService(
		redPacketManager,
	)
}

func GetRedPacketService() Service {
	return defaultServiceInstance
}
