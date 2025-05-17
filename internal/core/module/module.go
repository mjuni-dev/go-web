package module

import "github.com/mjuni.dev/go-web/internal/core/server/server"

type Config struct{}

type Module interface {
	Initialize() error
}

type RouteRegistry interface {
	RegisterRoutes(basePath string, registerFunc func(router server.Router))
}
