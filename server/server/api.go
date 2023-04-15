package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpackethandler "github.com/shengchaohua/red-packet-backend/server/handler/red_packet"
	"github.com/shengchaohua/red-packet-backend/server/route"
)

type apiServer struct {
	engine *gin.Engine
	config *config.ServerConfig
}

func NewApiServer() Server {
	serverConfig := config.GetGlobalConfig().ServerConfig
	if serverConfig == nil {
		panic("server config is nil")
	}
	if !serverConfig.IsApi() {
		panic("server config role is not api")
	}

	if serverConfig.IsDevEnv() {
		gin.SetMode(gin.TestMode)
	} else if serverConfig.IsLiveEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &apiServer{
		engine: gin.New(),
		config: serverConfig,
	}

	server.setLogger()
	server.registerHandler()

	return server
}

func (server *apiServer) Run(addr, port string) {
	if port != "" {
		run(server.engine, addr, port)
	} else {
		run(server.engine, server.config.Addr, server.config.Port)
	}
}

func (server *apiServer) setLogger() {
	setLogger(server.engine, server.config.LogFile)
}

func (server *apiServer) registerHandler() {
	redPacketHandler := redpackethandler.Handler{}
	server.engine.POST(
		route.ApiRouteCreateRedPacket,
		redPacketHandler.CreateRedPacket,
	)
	server.engine.POST(
		route.ApiRouteOpenRedPacket,
		redPacketHandler.OpenRedPacket,
	)
}
