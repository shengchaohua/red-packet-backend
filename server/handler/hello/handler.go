package hellohandler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	serverutils "github.com/shengchaohua/red-packet-backend/server/utils"
)

type Handler struct{}

func (handler *Handler) Hello(c *gin.Context) {
	var ctx = logger.NewCtxWithTraceId(context.Background(), "Hello")
	logger.Logger(ctx).Info("[Hello]start")

	c.JSON(http.StatusOK, serverutils.OkResponse(
		"Hello",
	))

	return
}
