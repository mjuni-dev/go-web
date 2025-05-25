package auth

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/features/auth/register"
	"github.com/mjuni.dev/go-web/web/features/auth/signin"
)

const MOD_WEB_AUTH string = "WEB.AUTH"

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
	return MOD_WEB_AUTH
}

func (m *Module) Initialize() error {
	m.registry.Register(signin.New(m.router))
	m.registry.Register(register.New(m.router))
	return nil
}
