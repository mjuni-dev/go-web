package server

import (
	"github.com/mjuni.dev/go-web/internal/core/server/echo"
	"github.com/mjuni.dev/go-web/internal/core/server/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/interfaces"
)

type ServerFactory struct{}

func New() *ServerFactory {
	return &ServerFactory{}
}

func (f *ServerFactory) EchoServer() interfaces.Server {
	return echo.New()
}

func (f *ServerFactory) GinServer() interfaces.Server {
	return gin.New()
}
