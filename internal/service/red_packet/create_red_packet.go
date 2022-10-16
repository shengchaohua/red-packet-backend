package redpacketservice

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
)

// CreateRedPacketRequest defines the request
// RedPacketName - optional
// Quantity - the max quantity of people that can open the red packet
// Amount - the money mount in the red packet
type CreateRedPacketRequest struct {
	RequestId         string                 `json:"request_id,omitempty"`
	UserId            uint64                 `json:"user_id,omitempty"`
	RedPacketCategory enum.RedPacketCategory `json:"red_packet_category,omitempty"`
	RedPacketType     enum.RedPacketType     `json:"red_packet_type,omitempty"`
	RedPacketName     string                 `json:"red_packet_name,omitempty"`
	Quantity          uint32                 `json:"quantity,omitempty"`
	Amount            uint32                 `json:"amount,omitempty"`
}

// CreateRedPacketResponse defines the reponse
type CreateRedPacketResponse struct {
	RequestId   string `json:"request_id,omitempty"`
	RedPacketId uint64 `json:"red_packet_id,omitempty"`
}

func (request *CreateRedPacketRequest) Validate(ctx context.Context) error {
	logger.Logger(ctx).Info("[CreateRedPacketRequest.Validte]start", zap.Any("request", request))

	if request.RedPacketCategory == 0 {
		return ErrWrongParam.WithMsg("red packet category is empty")
	}
	if request.RedPacketType == 0 {
		return ErrWrongParam.WithMsg("red packet type is empty")
	}
	return nil
}

func (service *defaultService) CreateRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {
	logger.Logger(ctx).Info("[CreateRedPacket]start", zap.Any("request", request))

	// insert red packet
	// deduct user money
	// generate use wallet transaction

	return &CreateRedPacketResponse{}, nil
}
