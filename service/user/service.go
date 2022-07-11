package userservice

import "context"

type Service interface {
	ClaimRedPacket(
		ctx context.Context,
		request *ClaimRedPacketRequest,
	) (*ClaimRedPacketResponse, error)
}

func InitService() {
	defaultService = &DefaultService{}
}

func GetService() Service {
	return defaultService
}
