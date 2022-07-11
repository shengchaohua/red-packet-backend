package redpackethanlder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	redpacketservice "github.com/shengchaohua/red-packet-backend/service/red_packet"
)

func CreateRedPacketHandler(ctx *gin.Context) {
	request := &redpacketservice.CreateRedPacketRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err_code": "0",
		})
	}

	newCtx := context.Background()
	response, err := redpacketservice.GetService().CreateRedPacket(newCtx, request)
	if err != nil {
		ctx.JSON(
			http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
