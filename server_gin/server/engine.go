package server

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

func newEngine(serverConfig *config.ServerConfig) *gin.Engine {
	if serverConfig == nil {
		panic("server config is nil")
	}

	if serverConfig.IsTestEnv() {
		gin.SetMode(gin.TestMode)
	} else if serverConfig.IsLiveEnv() {
		gin.SetMode(gin.ReleaseMode)
	}

	return gin.New()
}

func setLogger(engine *gin.Engine, logFile string) {
	if engine == nil {
		panic("engine is nil")
	}

	if logFile == "" {
		logFile = "./log/server.log"
		log.Println("log file is not specified, use default:", logFile)
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot open log file (%s): %w", logFile, err))
	}
	gin.DefaultWriter = io.MultiWriter(f)

	engine.Use(middleware.GetGinLogger())
}

func run(engine *gin.Engine, addr string) {
	if engine == nil || addr == "" {
		panic("engine is nil or addr is empty")
	}

	log.Println("Server is listening to", addr)
	httpServer := http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Server error:", err)
	}
	log.Println("Server exits")
}
