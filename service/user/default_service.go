package userservice

import (
	"context"
)

type DefaultService struct {
}

var defaultService *DefaultService

func (service *DefaultService) ClaimRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {
	return &CreateRedPacketResponse{}, nil
}
