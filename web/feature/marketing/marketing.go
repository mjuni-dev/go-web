package marketing

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/feature/marketing/about"
	"github.com/mjuni.dev/go-web/web/feature/marketing/landing"
	"github.com/mjuni.dev/go-web/web/feature/marketing/pricing"
)

const MOD_WEB_MARKETING string = "WEB.MARKETING"

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
	return MOD_WEB_MARKETING
}

func (m *Module) Initialize() error {
	m.registry.Register(landing.New(m.router))
	m.registry.Register(about.New(m.router))
	m.registry.Register(pricing.New(m.router))
	return nil
}
