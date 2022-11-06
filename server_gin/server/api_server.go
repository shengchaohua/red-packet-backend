package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpackethandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/red_packet"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
)

type apiServer struct {
	engine       *gin.Engine
	serverConfig *config.ServerConfig
}

func NewAPIServer() *apiServer {
	serverConfig := config.GetGlobalAppConfig().ServerConfig
	if serverConfig == nil {
		log.Fatalln("server config is nil")
	}
	if !serverConfig.IsAPI() {
		log.Fatalln("role is not api")
	}
	server := &apiServer{
		engine:       newEngine(serverConfig),
		serverConfig: serverConfig,
	}

	server.setLogger()
	server.registerHandler()

	return server
}

func (server *apiServer) Run() {
	run(server.engine, server.serverConfig.Addr)
}

func (server *apiServer) setLogger() {
	setLogger(server.engine, server.serverConfig.Log)
}

func (server *apiServer) registerHandler() {
	server.engine.POST(
		route.RouteCreateRedPacket,
		redpackethandler.CreateRedPacketHandler,
	)
	server.engine.POST(
		route.RouteOpenRedPacket,
		redpackethandler.OpenRedPacketHandler,
	)
}
