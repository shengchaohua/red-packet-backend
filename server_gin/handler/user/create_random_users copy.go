package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	userservice "github.com/shengchaohua/red-packet-backend/internal/service/user"
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
	serverutils "github.com/shengchaohua/red-packet-backend/server_gin/utils"
	"go.uber.org/zap"
)

func TopupUserWalletHandler(c *gin.Context) {
	var request = &userservice.CreateRandomUsersRequest{}
	if c.ShouldBindJSON(request) != nil {
		c.JSON(http.StatusOK, serverutils.BadRequest())
		return
	}

	var (
		ctx      = logger.NewCtxWithTraceId()
		response *userservice.CreateRandomUsersResponse
		err      error
	)
	logger.Logger(ctx).Info("[CreateRandomUsersHandler]start", zap.Any("request", request))

	if err = request.Validate(); err != nil {
		logger.Logger(ctx).Error("[CreateRedPacketHandler]validate_request_error", zap.Error(err))
		c.JSON(http.StatusOK, serverutils.WrongParam(errorgrouppkg.GetErrmsg(err)))
		return
	}

	response, err = userservice.GetUserService().CreateRandomUsers(ctx, request)
	if err != nil {
		logger.Logger(ctx).Error("[CreateRandomUsersHandler]service_error", zap.Error(err))
		c.JSON(http.StatusOK, serverutils.ServerError())
		return
	}

	logger.Logger(ctx).Info("[CreateRandomUsersHandler]response", zap.Any("response", response))
	c.JSON(http.StatusOK, serverutils.OkResponse(response))

	return
}
