package redpackethandler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"
	serverutils "github.com/shengchaohua/red-packet-backend/server/utils"
	"go.uber.org/zap"
)

var (
	createRedPacketErrors = map[int]bool{
		constants.Errcode_WalletBalanceNotEnough: true,
		constants.Errcode_UserNotInGroup:         true,
	}
)

func (handler *Handler) CreateRedPacket(c *gin.Context) {
	var request = &redpacketservice.CreateRedPacketRequest{}
	if c.ShouldBindJSON(request) != nil {
		c.JSON(http.StatusOK, serverutils.BadRequest())
		return
	}

	var (
		ctx      = logger.NewCtxWithTraceId(context.Background(), "CreateRedPacket")
		response *redpacketservice.CreateRedPacketResponse
		err      error
	)
	logger.Logger(ctx).Info("[CreateRedPacket]start", zap.Any("request", request))

	if err = request.Validate(); err != nil {
		logger.Logger(ctx).Error("[CreateRedPacket]validate_request_error", zap.Error(err))
		c.JSON(http.StatusOK, serverutils.WrongParam(errorpkg.GetErrmsg(err)))
		return
	}

	response, err = redpacketservice.GetRedPacketService().CreateRedPacket(ctx, request)
	if err != nil {
		logger.Logger(ctx).Error("[CreateRedPacket]service_error", zap.Error(err))
		errcode, ok := errorpkg.GetErrcode(err)
		if ok && createRedPacketErrors[errcode] {
			c.JSON(http.StatusOK, serverutils.Response(
				errcode,
				errorpkg.GetErrmsg(err),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, serverutils.ServerError())
		return
	}

	logger.Logger(ctx).Info("[CreateRedPacket]response", zap.Any("response", response))
	c.JSON(http.StatusOK, serverutils.OkResponse(response))

	return
}
