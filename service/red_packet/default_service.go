package redpacketservice

import (
	"context"

	redpacketpkg "github.com/shengchaohua/red-packet-backend/data/pkg/red_packet"
)

type defaultService struct {
	redPacketManager redpacketpkg.Manager
}

func NewDefaultService(
	redPacketManager redpacketpkg.Manager,
) Service {
	return &defaultService{
		redPacketManager: redPacketManager,
	}
}

func (service *defaultService) CreateRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {

	// insert red packet
	// deduct user money
	// generate use wallet transaction

	return &CreateRedPacketResponse{}, nil
}
