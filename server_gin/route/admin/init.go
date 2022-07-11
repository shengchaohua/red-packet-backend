package adminroute

import (
	"github.com/gin-gonic/gin"

	redpackethanlder "github.com/shengchaohua/red-packet-backend/server_gin/handler/admin/red_packet"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST(RouteCreateRedPacket, redpackethanlder.CreateRedPacketHandler)

	return router
}
