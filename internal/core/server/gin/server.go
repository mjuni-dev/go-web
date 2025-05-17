package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/interfaces"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) Start(address string) error {
	return s.engine.Run(address)
}

func (s *Server) Router() interfaces.Router {
	return &Router{
		gin: s.engine,
	}
}
