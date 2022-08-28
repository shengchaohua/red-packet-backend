package apiserver

import (
	"github.com/gin-gonic/gin"

	redpackethandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/red_packet"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	router := &Server{
		Engine: gin.Default(),
	}

	router.RegisterHandler()

	return router
}

func (server *Server) RegisterHandler() {
	server.Engine.POST(
		redpackethandler.RouteCreateRedPacket,
		redpackethandler.CreateRedPacketHandler,
	)
	server.Engine.POST(
		redpackethandler.RouteOpenRedPacket,
		redpackethandler.OpenRedPacketHandler,
	)
}

func (server *Server) Run() {
	if err := server.Engine.Run(); err != nil {
		panic(err)
	}
}
