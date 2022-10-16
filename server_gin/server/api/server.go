package apiserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpackethandler "github.com/shengchaohua/red-packet-backend/server_gin/handler/red_packet"
	"github.com/shengchaohua/red-packet-backend/server_gin/middleware"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
)

type apiServer struct {
	engine *gin.Engine
}

func NewServer() *apiServer {
	server := newServer()
	server.setLogger()
	server.registerHandler()
	return server
}

func newServer() *apiServer {
	apiConfig := config.GetGlobalAppConfig().APIConfig
	if apiConfig == nil {
		panic("APIConfig is nil")
	}

	if apiConfig.IsTestEnv() {
		gin.SetMode(gin.TestMode)
	} else if apiConfig.IsLiveEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	return &apiServer{
		engine: gin.New(),
	}
}

func (server *apiServer) setLogger() {
	gin.DisableConsoleColor()

	logFile := config.GetGlobalAppConfig().APIConfig.Log
	if logFile == "" {
		logFile = "./log/api_server.log"
		log.Println("log file is not specified, use default:", logFile)
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot open log file (%s): %w", logFile, err))
	}
	gin.DefaultWriter = io.MultiWriter(f)

	server.engine.Use(middleware.GetGinLogger())
}

func (server *apiServer) registerHandler() {
	server.engine.POST(
		route.RouteOpenRedPacket,
		redpackethandler.OpenRedPacketHandler,
	)
}

func (server *apiServer) Run() {
	addr := config.GetGlobalAppConfig().APIConfig.Addr
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
