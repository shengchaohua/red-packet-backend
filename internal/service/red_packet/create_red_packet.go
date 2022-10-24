package redpacketservice

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

// CreateRedPacketRequest defines the request
// RedPacketName - optional
// Quantity - the max quantity of people that can open the red packet
// Amount - the money amount in the red packet
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

	if request.RequestId == "" {
		return ErrWrongParam.WithMsg("request id is empty")
	}
	if request.UserId == 0 {
		return ErrWrongParam.WithMsg("user id is empty")
	}
	if request.RedPacketCategory == 0 {
		return ErrWrongParam.WithMsg("red packet category is empty")
	}
	if request.RedPacketType == 0 {
		return ErrWrongParam.WithMsg("red packet type is empty")
	}
	if request.Quantity == 0 {
		return ErrWrongParam.WithMsg("red packet quantity is zero")
	}
	if request.Amount == 0 {
		return ErrWrongParam.WithMsg("red packet amount is zero")
	}

	return nil
}

func (service *defaultService) CreateRedPacket(
	ctx context.Context,
	request *CreateRedPacketRequest,
) (*CreateRedPacketResponse, error) {
	logger.Logger(ctx).Info("[defaultService.CreateRedPacket]start", zap.Any("request", request))

	var (
		redPacket *redpacketmodel.RedPacket
		err       error
	)
	_, err = service.EngineManager.GetMasterEngine().Transaction(func(session *xorm.Session) (interface{}, error) {
		if redPacket, err = service.redPacketManager.CreateRedPacket(
			ctx,
			session,
			request.RedPacketName,
			request.RedPacketCategory,
			request.RedPacketType,
			request.Quantity,
			request.Amount,
		); err != nil {
			return nil, err
		}

		if err := service.userWalletManager.DeductUserWallet(
			ctx,
			session,
			request.UserId,
			redPacket.Amount,
		); err != nil {
			return nil, err
		}

		return redPacket, nil
	})
	if err != nil {
		return nil, ErrServer.WrapWithMsg(err, "[CreateRedPacket]transaction_error")
	}

	return &CreateRedPacketResponse{
		RequestId:   request.RequestId,
		RedPacketId: redPacket.Id,
	}, nil
}
