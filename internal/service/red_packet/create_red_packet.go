package redpacketservice

import (
	"context"
	"fmt"

	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

// CreateRedPacketRequest defines the request
// - RedPacketName: optional
// - Quantity: the max quantity of people that can open the red packet
// - Amount: the money amount
// - ReceiverUserId: only valid when red packet category is P2P
// - GroupId: only valid when
type CreateRedPacketRequest struct {
	RequestId           string                   `json:"request_id,omitempty"`
	UserId              uint64                   `json:"user_id,omitempty"`
	RedPacketCategory   enum.RedPacketCategory   `json:"red_packet_category,omitempty"`
	RedPacketResultType enum.RedPacketResultType `json:"red_packet_type,omitempty"`
	RedPacketName       string                   `json:"red_packet_name,omitempty"`
	Quantity            uint32                   `json:"quantity,omitempty"`
	Amount              uint32                   `json:"amount,omitempty"`
	ReceiverUserId      uint64                   `json:"receiver_user_id,omitempty"`
	GroupId             uint64                   `json:"group_id,omitempty"`
}

// CreateRedPacketResponse defines the reponse
type CreateRedPacketResponse struct {
	RequestId   string `json:"request_id,omitempty"`
	RedPacketId uint64 `json:"red_packet_id,omitempty"`
}

func (request *CreateRedPacketRequest) Validate() error {
	if request.RequestId == "" {
		return ErrWrongParam.WithMsg("request id is empty")
	}
	if request.UserId == 0 {
		return ErrWrongParam.WithMsg("user id is empty")
	}
	if request.RedPacketCategory == 0 {
		return ErrWrongParam.WithMsg("red packet category is empty")
	}
	if request.Amount == 0 {
		return ErrWrongParam.WithMsg("red packet amount is zero")
	}
	switch request.RedPacketCategory {
	case enum.RedPacketCategoryP2P:
		if request.Quantity != 1 {
			return ErrWrongParam.WithMsg("red packet quantity is not one")
		}
		if request.ReceiverUserId == 0 {
			return ErrWrongParam.WithMsg("red packet receiver user id is empty")
		}
		if request.ReceiverUserId == request.UserId {
			return ErrWrongParam.WithMsg("cannot send red packet to yourself")
		}
	case enum.RedPacketCategoryGroup:
		if request.RedPacketResultType == 0 {
			return ErrWrongParam.WithMsg("red packet type is empty")
		}
		if request.Quantity == 0 {
			return ErrWrongParam.WithMsg("red packet quantity is empty")
		}
		if request.GroupId == 0 {
			return ErrWrongParam.WithMsg("red packet group id is empty")
		}
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
	_, err = service.GetMasterEngine().Transaction(func(session *xorm.Session) (interface{}, error) {
		// create red packet
		switch request.RedPacketCategory {
		case enum.RedPacketCategoryP2P:
			redPacket, err = service.createP2PRedPacket(ctx, session, request)
			if err != nil {
				return nil, err
			}
		case enum.RedPacketCategoryGroup:
			redPacket, err = service.createGroupRedPacket(ctx, session, request)
			if err != nil {
				return nil, err
			}
		}

		// deduct money
		amount := request.Amount
		if request.RedPacketResultType == enum.RedPacketResultTypeIdenticalAmount {
			amount = redPacket.Amount * redPacket.Quantity
		}
		err = service.userWalletManager.DeductUserWalletBalance(ctx, session, request.UserId, amount)
		if err != nil {
			return nil, err
		}

		// save user wallet txn
		err = service.redPacketTxnManager.AddUserWalletTxn(
			ctx,
			session,
			request.UserId,
			enum.CreateRedPacket,
			fmt.Sprint(redPacket.Id),
			amount,
		)
		if err != nil {
			return nil, err
		}

		return redPacket, nil
	})
	if err != nil {
		logger.Logger(ctx).Error("[defaultService.CreateRedPacket]transaction_error", zap.Any("error", err))
		return nil, errorMapping(err)
	}

	return &CreateRedPacketResponse{
		RequestId:   request.RequestId,
		RedPacketId: redPacket.Id,
	}, nil
}

func (service *defaultService) createP2PRedPacket(
	ctx context.Context,
	session *xorm.Session,
	request *CreateRedPacketRequest,
) (*redpacketmodel.RedPacket, error) {
	// TODO check user relation
	return service.redPacketManager.CreateP2PRedPacket(
		ctx,
		session,
		request.RedPacketName,
		request.RedPacketResultType,
		request.Quantity,
		request.Amount,
		request.ReceiverUserId,
	)
}

func (service *defaultService) createGroupRedPacket(
	ctx context.Context,
	session *xorm.Session,
	request *CreateRedPacketRequest,
) (*redpacketmodel.RedPacket, error) {
	userInGroup, err := service.userGroupRelationManager.CheckUserInGroup(
		ctx,
		request.UserId,
		request.GroupId,
	)
	if err != nil {
		return nil, ErrServer.WrapWithMsg(err, "check_user_in_group_error")
	}
	if !userInGroup {
		return nil, ErrUserNotInGroup.WithMsg("user is not in this group")
	}

	return service.redPacketManager.CreateGroupRedPacket(
		ctx,
		session,
		request.RedPacketName,
		request.RedPacketResultType,
		request.Quantity,
		request.Amount,
		request.GroupId,
	)
}
