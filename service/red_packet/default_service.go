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
	// check
	if request.RedPacketCategory == 0 {
		return nil, ErrInvalidParams.WithMsg("[CreateRedPacket]red packet category is empty")
	}
	if request.RedPacketType == 0 {
		return nil, ErrInvalidParams.WithMsg("[CreateRedPacket]red packet type is empty")
	}

	//

	return &CreateRedPacketResponse{}, nil
}
