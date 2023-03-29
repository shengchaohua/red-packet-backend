package redpackethandler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/shengchaohua/red-packet-backend/internal/constants"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
	serverutils "github.com/shengchaohua/red-packet-backend/server_gin/utils"
)

func CreateRedPacketHandler(c *gin.Context) {
	var request = &redpacketservice.CreateRedPacketRequest{}
	if c.ShouldBindJSON(request) != nil {
		c.JSON(http.StatusOK, serverutils.BadRequest())
		return
	}

	var (
		ctx      = logger.NewCtxWithTraceId(context.Background(), "CreateRedPacketHandler")
		response *redpacketservice.CreateRedPacketResponse
		err      error
	)
	logger.Logger(ctx).Info("[CreateRedPacketHandler]start", zap.Any("request", request))

	if err = request.Validate(); err != nil {
		logger.Logger(ctx).Error("[CreateRedPacketHandler]validate_request_error", zap.Error(err))
		c.JSON(http.StatusOK, serverutils.WrongParam(errorgrouppkg.GetErrmsg(err)))
		return
	}

	response, err = redpacketservice.GetRedPacketService().CreateRedPacket(ctx, request)
	if err != nil {
		logger.Logger(ctx).Error("[CreateRedPacketHandler]service_error", zap.Error(err))
		errcode, ok := errorgrouppkg.GetErrcode(err)
		if ok && allowedErrorMap[route.RouteCreateRedPacket][constants.Errcode(errcode)] {
			c.JSON(http.StatusOK, serverutils.Response(
				errcode,
				errorgrouppkg.GetErrmsg(err),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, serverutils.ServerError())
		return
	}

	logger.Logger(ctx).Info("[CreateRedPacketHandler]response", zap.Any("response", response))
	c.JSON(http.StatusOK, serverutils.OkResponse(response))

	return
}
