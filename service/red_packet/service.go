package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/data/pkg/red_packet"
)

type Service interface {
	CreateRedPacket(ctx context.Context, request *CreateRedPacketRequest) (*CreateRedPacketResponse, error)
}

var defaultServiceInstance Service

func InitService() {
	redPacketManager := redpacketpkg.GetDefaultManager()
	defaultServiceInstance = NewDefaultService(
		redPacketManager,
	)
}

func GetDefaultService() Service {
	return defaultServiceInstance
}
