package redpacketservice

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
)

// OpenRedPacketRequest defines the request
type OpenRedPacketRequest struct {
	RequestId   string `json:"request_id,omitempty"`
	UserId      uint64 `json:"user_id,omitempty"`
	RedPacketId uint32 `json:"red_packet_id,omitempty"`
}

// CreateRedPacketResponse defines the reponse
type OpenRedPacketResponse struct {
	RequestId    string `json:"request_id,omitempty"`
	ResultAmount uint64 `json:"result_amount,omitempty"`
}

func (request *OpenRedPacketRequest) Validate() error {
	if request.RequestId == "" {
		return ErrWrongParam.WithMsg("request id is empty")
	}
	if request.UserId == 0 {
		return ErrWrongParam.WithMsg("user id is empty")
	}
	if request.RedPacketId == 0 {
		return ErrWrongParam.WithMsg("red packet id is empty")
	}
	return nil
}

func (service *defaultService) OpenRedPacket(
	ctx context.Context,
	request *OpenRedPacketRequest,
) (*OpenRedPacketRequest, error) {
	logger.Logger(ctx).Info("[defaultService.OpenRedPacket]start", zap.Any("request", request))

	return &OpenRedPacketRequest{
		RequestId: request.RequestId,
	}, nil
}
