package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

type Server struct {
	engine *echo.Echo
}

func New() *Server {
	return &Server{
		engine: echo.New(),
	}
}

func (s *Server) Start(address string) error {
	return s.engine.Start(address)
}

func (s *Server) Router() server.Router {
	return &Router{
		echo: s.engine,
	}
}
