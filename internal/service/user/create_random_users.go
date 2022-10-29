package userservice

import "context"

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
	return &CreateRandomUsersResponse{}, nil
}
