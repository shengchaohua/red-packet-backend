package redpackethandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		c.JSON(http.StatusOK, serverutils.WrongParam(errorgrouppkg.GetErrmsg(err)))
		return
	}

	response, err := redpacketservice.GetRedPacketService().CreateRedPacket(ctx, request)
	if err != nil {
		c.JSON(http.StatusOK, serverutils.ServerError())
		return
	}

	c.JSON(http.StatusOK, serverutils.OkResponse(response))
	return
}
