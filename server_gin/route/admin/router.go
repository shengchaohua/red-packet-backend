package adminroute

import (
	"github.com/gin-gonic/gin"

	redpackethandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/red_packet"
	"github.com/shengchaohua/red-packet-backend/server_gin/route/routes"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: gin.Default(),
	}

	router.RegisterHandler()

	return router
}

func (router *Router) RegisterHandler() {
	router.POST(routes.RouteOpenRedPacket, redpackethandler.CreateRedPacketHandler)
}

func (router *Router) Run() {
	if err := router.Engine.Run(); err != nil {
		panic(err)
	}
}
