package userservice

import "context"

type Service interface {
	OpenRedPacket(ctx context.Context, request *OpenRedPacketRequest) (*OpenRedPacketResponse, error)
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
