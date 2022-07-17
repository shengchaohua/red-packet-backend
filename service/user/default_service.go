package userservice

import (
	"context"
)

type defaultService struct {
}

func NewDefaultService() *defaultService {
	return &defaultService{}
}

func (service *defaultService) OpenRedPacket(
	ctx context.Context,
	request *OpenRedPacketRequest,
) (*OpenRedPacketResponse, error) {
	return &OpenRedPacketResponse{}, nil
}
