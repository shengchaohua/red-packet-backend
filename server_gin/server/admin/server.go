package adminserver

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	server := &Server{
		Engine: gin.New(),
	}

	server.RegisterHandler()

	return server
}

func (server *Server) RegisterHandler() {

}

func (server *Server) Run() {
	if err := server.Engine.Run(); err != nil {
		panic(err)
	}
}
