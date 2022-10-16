package adminserver

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpackethandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/red_packet"
	"github.com/shengchaohua/red-packet-backend/server_gin/routes"
)

type adminServer struct {
	*gin.Engine
}

func NewServer() *adminServer {
	setGinMode()
	server := &adminServer{
		Engine: gin.New(),
	}

	server.setLogger()
	server.RegisterHandler()
	server.Use(gin.Recovery())

	return server
}

func setGinMode() {
	adminServerConfig := config.GetGlobalAppConfig().AdminConfig
	if adminServerConfig.IsTestEnv() {
		gin.SetMode(gin.TestMode)
	} else if adminServerConfig.IsLiveEnv() {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (server *adminServer) setLogger() {
	gin.DisableConsoleColor()

	logFile := config.GetGlobalAppConfig().AdminConfig.Log
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}

func (server *adminServer) RegisterHandler() {
	server.Engine.POST(
		routes.RouteCreateRedPacket,
		redpackethandler.CreateRedPacketHandler,
	)
}

func (server *adminServer) Run() {
	adminServerConfig := config.GetGlobalAppConfig().AdminConfig
	if err := server.Engine.Run(adminServerConfig.Addr); err != nil {
		panic(fmt.Errorf("run admin server error: %w", err))
	}
}
