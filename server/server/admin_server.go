package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	userhandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/user"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
)

type adminServer struct {
	engine       *gin.Engine
	serverConfig *config.ServerConfig
}

func NewAdminServer() *adminServer {
	serverConfig := config.GetGlobalAppConfig().ServerConfig
	if serverConfig == nil {
		log.Fatalln("server config is nil")
	}
	if !serverConfig.IsAdmin() {
		log.Fatalln("role is not admin")
	}
	server := &adminServer{
		engine:       newEngine(serverConfig),
		serverConfig: serverConfig,
	}

	server.setLogger()
	server.registerHandler()

	return server
}

func (server *adminServer) Run() {
	run(server.engine, server.serverConfig.Addr)
}

func (server *adminServer) setLogger() {
	setLogger(server.engine, server.serverConfig.Log)
}

func (server *adminServer) registerHandler() {
	server.engine.POST(
		route.RouteCreateRandomUsers,
		userhandler.CreateRandomUsersHandler,
	)
}
