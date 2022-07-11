package redpacketservice

import "context"

type Service interface {
	CreateRedPacket(ctx context.Context, request *CreateRedPacketRequest) (*CreateRedPacketResponse, error)
}

var (
	defaultService Service
)

func InitService() {
	defaultService = &DefaultService{}
}

func GetService() Service {
	return defaultService
}
