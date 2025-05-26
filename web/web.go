package web

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/feature/auth"
	"github.com/mjuni.dev/go-web/web/feature/marketing"
)

const MOD_WEB string = "WEB"

type Module struct {
	registry *module.Registry
	router   server.Router
}

func New(registry *module.Registry, router server.Router) *Module {
	return &Module{
		registry: registry,
		router:   router,
	}
}

func (m *Module) Name() string {
	return MOD_WEB
}

func (m *Module) Initialize() error {
	m.registry.Register(auth.New(m.registry, m.router))
	m.registry.Register(marketing.New(m.registry, m.router))

	return nil
}
