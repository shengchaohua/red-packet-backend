package redpacketservice

import (
	"context"
)

type DefaultService struct {
}

func (service *DefaultService) CreateRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {
	return &CreateRedPacketResponse{}, nil
}
