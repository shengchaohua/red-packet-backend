package redpackethandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	userservice "github.com/shengchaohua/red-packet-backend/internal/service/user"
)

func OpenRedPacketHandler(ctx *gin.Context) {
	request := &userservice.OpenRedPacketRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err_code": "0",
		})
	}

	newCtx := context.Background()
	response, err := userservice.GetDefaultService().OpenRedPacket(newCtx, request)
	if err != nil {
		ctx.JSON(
			http.StatusOK, gin.H{})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
