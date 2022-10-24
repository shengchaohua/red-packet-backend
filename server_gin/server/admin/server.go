package adminserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	"github.com/shengchaohua/red-packet-backend/server_gin/middleware"
)

type adminServer struct {
	engine *gin.Engine
}

func NewServer() *adminServer {
	server := newServer()
	server.setLogger()
	server.registerHandler()
	return server
}

func newServer() *adminServer {
	adminConfig := config.GetGlobalAppConfig().AdminConfig
	if adminConfig == nil {
		panic("AdminConfig is nil")
	}

	if adminConfig.IsTestEnv() {
		gin.SetMode(gin.TestMode)
	} else if adminConfig.IsLiveEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	return &adminServer{
		engine: gin.New(),
	}
}

func (server *adminServer) setLogger() {
	gin.DisableConsoleColor()

	logFile := config.GetGlobalAppConfig().AdminConfig.Log
	if logFile == "" {
		logFile = "./log/admin_server.log"
		log.Println("log file is not specified, use default:", logFile)
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot open log file (%s): %w", logFile, err))
	}
	gin.DefaultWriter = io.MultiWriter(f)

	server.engine.Use(middleware.GetGinLogger())
}

func (server *adminServer) registerHandler() {

}

func (server *adminServer) Run() {
	addr := config.GetGlobalAppConfig().AdminConfig.Addr
	if addr == "" {
		panic("server address is empty")
	}

	log.Println("Server is listening to", addr)
	httpServer := http.Server{
		Addr:    addr,
		Handler: server.engine.Handler(),
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Server error:", err)
	}
}
