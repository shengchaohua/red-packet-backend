package redpacketservice

import (
	"context"

	redpacketagent "github.com/shengchaohua/red-packet-backend/data/agent/red_packet"
)

type defaultService struct {
	redPacketAgent redpacketagent.Agent
}

func NewDefaultService(
	redPacketAgent redpacketagent.Agent,
) Service {
	return &defaultService{
		redPacketAgent: redPacketAgent,
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
