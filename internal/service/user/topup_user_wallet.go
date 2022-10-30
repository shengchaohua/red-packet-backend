package userservice

import (
	"context"

	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"go.uber.org/zap"
)

// TopupUserWalletRequest defines request
type TopupUserWalletRequest struct {
	RequestId string `json:"request_id,omitempty"`
	UserId    uint64 `json:"user_id,omitempty"`
	Amount    uint64 `json:"amount,omitempty"`
}

type TopupUserWalletResponse struct {
	RequestId string `json:"request_id,omitempty"`
}

func (request *TopupUserWalletRequest) Validate() error {
	if request.RequestId == "" {
		return ErrWrongParam.WithMsg("request id is empty")
	}
	if request.Amount == 0 {
		return ErrWrongParam.WithMsg("topup amount is empty")
	}
	return nil
}

func (service *defaultService) TopupUserWallet(
	ctx context.Context,
	request *TopupUserWalletRequest,
) (*TopupUserWalletResponse, error) {
	logger.Logger(ctx).Info("[defaultService.TopupUserWallet]start", zap.Any("request", request))

	return &TopupUserWalletResponse{
		RequestId: request.RequestId,
	}, nil
}
