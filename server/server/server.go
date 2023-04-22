package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shengchaohua/red-packet-backend/server/middleware"
)

type Server interface {
	Run(port string)
}

func setLogger(engine *gin.Engine, logFile string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot open log file (%s): %w", logFile, err))
	}
	gin.DefaultWriter = io.MultiWriter(f)

	engine.Use(middleware.GetGinLogger())
}

func run(engine *gin.Engine, port string) {
	var fullAddr = fmt.Sprintf(":%s", port)
	log.Println("Server is listening to", fullAddr)

	httpServer := http.Server{
		Addr:    fullAddr,
		Handler: engine.Handler(),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Server error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // blocking
	log.Println("Server receives exit signal")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
	log.Println("Server exits")
}
