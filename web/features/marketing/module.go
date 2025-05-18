package marketing

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
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

	return nil
}
