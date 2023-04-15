package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	hellohandler "github.com/shengchaohua/red-packet-backend/server/handler/hello"
	userhandler "github.com/shengchaohua/red-packet-backend/server/handler/user"
	"github.com/shengchaohua/red-packet-backend/server/route"
)

type adminServer struct {
	engine *gin.Engine
	config *config.ServerConfig
}

func NewAdminServer() Server {
	serverConfig := config.GetGlobalConfig().ServerConfig
	if serverConfig == nil {
		panic("server config is nil")
	}
	if !serverConfig.IsAdmin() {
		panic("server config role is not admin")
	}

	//if serverConfig.IsDevEnv() {
	//	gin.SetMode(gin.DebugMode)
	//} else if serverConfig.IsLiveEnv() {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	server := &adminServer{
		engine: gin.New(),
		config: serverConfig,
	}

	server.registerHandler()

	return server
}

func (server *adminServer) Run(addr, port string) {
	if addr != "" && port != "" {
		run(server.engine, addr, port)
	} else {
		run(server.engine, server.config.Addr, server.config.Port)
	}
}

func (server *adminServer) setLogger() {
	setLogger(server.engine, server.config.LogFile)
}

func (server *adminServer) registerHandler() {
	helloHandler := &hellohandler.Handler{}
	server.engine.GET(
		route.AdminRouteHello,
		helloHandler.Hello,
	)

	userHandler := &userhandler.Handler{}
	server.engine.POST(
		route.AdminRouteRegister,
		userHandler.Register,
	)
}
