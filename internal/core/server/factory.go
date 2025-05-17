package server

import (
	"github.com/mjuni.dev/go-web/internal/core/server/echo"
	"github.com/mjuni.dev/go-web/internal/core/server/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

type ServerFactory struct{}

func NewServerFactory() *ServerFactory {
	return &ServerFactory{}
}

func (f *ServerFactory) EchoServer() server.Server {
	return echo.New()
}

func (f *ServerFactory) GinServer() server.Server {
	return gin.New()
}
