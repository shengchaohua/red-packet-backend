package redpacketservice

import (
	"context"

	redpacketagent "github.com/shengchaohua/red-packet-backend/data/agent/red_packet"
)

type Service interface {
	CreateRedPacket(ctx context.Context, request *CreateRedPacketRequest) (*CreateRedPacketResponse, error)
}

var defaultServiceInstance Service

func InitService() {
	redPacketAgent := redpacketagent.GetDefaultAgent()
	defaultServiceInstance = NewDefaultService(
		redPacketAgent,
	)
}

func GetDefaultService() Service {
	return defaultServiceInstance
}
