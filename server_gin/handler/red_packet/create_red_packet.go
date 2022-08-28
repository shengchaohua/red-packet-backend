package redpackethandler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	serverutils "github.com/shengchaohua/red-packet-backend/server_gin/utils"
	redpacketservice "github.com/shengchaohua/red-packet-backend/service/red_packet"
)

func CreateRedPacketHandler(ginCtx *gin.Context) {
	var request = &redpacketservice.CreateRedPacketRequest{}
	if ginCtx.ShouldBindJSON(request) != nil {
		ginCtx.JSON(http.StatusOK, serverutils.BadRequest())
		return
	}

	// check param
	if request.RedPacketCategory == 0 {
		ginCtx.JSON(http.StatusOK, serverutils.WrongParam("red packet category is empty"))
		return
	}
	if request.RedPacketType == 0 {
		ginCtx.JSON(http.StatusOK, serverutils.WrongParam("red packet type is empty"))
		return
	}

	// service
	ctx := context.Background()
	response, err := redpacketservice.GetDefaultService().CreateRedPacket(ctx, request)
	if err != nil {
		// log error
		ginCtx.JSON(http.StatusOK, serverutils.ServerError())
		return
	}

	ginCtx.JSON(http.StatusOK, serverutils.WrapResponse(response))
	return
}
