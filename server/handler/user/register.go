package userhandler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	serverutils "github.com/shengchaohua/red-packet-backend/server/utils"
)

func (handler *Handler) Register(c *gin.Context) {
	var ctx = logger.NewCtxWithTraceId(context.Background(), "Register")
	logger.Logger(ctx).Info("[Hello]start")

	c.JSON(http.StatusOK, serverutils.OkResponse("SUCCESS"))

	return
}
