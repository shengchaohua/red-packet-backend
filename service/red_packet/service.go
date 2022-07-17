package redpacketservice

import "context"

type Service interface {
	CreateRedPacket(ctx context.Context, request *CreateRedPacketRequest) (*CreateRedPacketResponse, error)
}

var (
	defaultServiceInstance Service
)

func InitService() {
	defaultServiceInstance = NewDefaultService()
}

func GetDefaultService() Service {
	return defaultServiceInstance
}
