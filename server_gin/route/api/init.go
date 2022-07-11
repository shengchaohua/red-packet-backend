package apiroute

import (
	"github.com/gin-gonic/gin"

	userhandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/admin/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST(RouteClaimRedPacket, userhandler.ClaimRedPacketHandler)

	return router
}
