package redpacketservice

import (
	"context"
)

type defaultService struct {
}

func NewDefaultService() *defaultService {
	return &defaultService{}
}

func (service *defaultService) CreateRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {
	return &CreateRedPacketResponse{}, nil
}
