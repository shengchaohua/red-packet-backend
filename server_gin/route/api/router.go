package apiroute

import (
	"github.com/gin-gonic/gin"

	userhandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/admin/user"
	"github.com/shengchaohua/red-packet-backend/server_gin/route/routes"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func (router *Router) RegisterHandler() {
	router.POST(routes.RouteOpenRedPacket, userhandler.OpenRedPacketHandler)
}
