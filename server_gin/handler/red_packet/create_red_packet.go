package redpackethandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/shengchaohua/red-packet-backend/internal/constants"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
	serverutils "github.com/shengchaohua/red-packet-backend/server_gin/utils"
)

func CreateRedPacketHandler(c *gin.Context) {
	var request = &redpacketservice.CreateRedPacketRequest{}
	if c.ShouldBindJSON(request) != nil {
		c.JSON(http.StatusOK, serverutils.BadRequest())
		return
	}

	ctx := logger.NewCtxWithTraceId()
	if err := request.Validate(ctx); err != nil {
		logger.Logger(ctx).Error("[CreateRedPacketHandler]validate_request_error", zap.Error(err))
		c.JSON(http.StatusOK, serverutils.WrongParam(errorgrouppkg.GetErrmsg(err)))
		return
	}

	response, err := redpacketservice.GetService().CreateRedPacket(ctx, request)
	if err != nil {
		logger.Logger(ctx).Error("[CreateRedPacketHandler]create_red_packet_error", zap.Error(err))
		if errcode, ok := errorgrouppkg.GetErrcode(err); ok {
			errcodeEnum := constants.ParseErrcodeEnum(errcode)
			c.JSON(http.StatusOK, serverutils.Response(
				errcodeEnum,
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
