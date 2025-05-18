package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	e := gin.Default()
	e.Static("/public", "./web/public")
	return &Server{
		engine: e,
	}
}

func (s *Server) Start(address string) error {
	return s.engine.Run(address)
}

func (s *Server) Router() server.Router {
	return &Router{
		gin: s.engine,
	}
}
