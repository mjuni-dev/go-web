package web

import (
	"fmt"

	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/features/marketing/pricing"
)

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

func (m *Module) Initialize() error {
	fmt.Println(" >> Initializing WEB module...")
	m.registry.Register(pricing.New(m.router))

	return nil
}
