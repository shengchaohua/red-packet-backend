package userservice

import (
	"context"
	"fmt"

	usermodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"github.com/shengchaohua/red-packet-backend/internal/utils"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

// CreateRandomUsersRequest defines request
// - Count: the number of users that will be created
// - EnableWallet: whether to create user wallet or not
type CreateRandomUsersRequest struct {
	RequestId    string `json:"request_id,omitempty"`
	Count        uint32 `json:"count,omitempty"`
	EnableWallet bool   `json:"enable_wallet,omitempty"`
}

type CreateRandomUsersResponse struct {
	RequestId string `json:"request_id,omitempty"`
}

func (request *CreateRandomUsersRequest) Validate() error {
	if request.RequestId == "" {
		return ErrWrongParam.WithMsg("request id is empty")
	}
	if request.Count == 0 {
		return ErrWrongParam.WithMsg("count is empty")
	}
	return nil
}

func (service *defaultService) CreateRandomUsers(
	ctx context.Context,
	request *CreateRandomUsersRequest,
) (*CreateRandomUsersResponse, error) {
	logger.Logger(ctx).Info("[defaultService.CreateRandomUsers]start", zap.Any("request", request))

	for idx := 0; idx < int(request.Count); idx++ {
		var (
			user *usermodel.User
			err  error
		)
		_, err = service.GetMasterEngine().Transaction(func(session *xorm.Session) (interface{}, error) {
			user, err = service.userManager.CreateUser(
				ctx,
				session,
				fmt.Sprintf("TestUser-%d-%d", utils.GetCurrentTime(), idx),
			)
			if err != nil {
				return nil, err
			}

			if request.EnableWallet {
				err = service.useWalletManager.CreateUserWallet(ctx, session, user.Id)
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		})
		if err != nil {
			logger.Logger(ctx).Error("[defaultService.CreateRandomUsers]transaction_error", zap.Any("error", err))
			return nil, err
		}
	}

	return &CreateRandomUsersResponse{}, nil
}
