package feature

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/feature/auth"
	"github.com/mjuni.dev/go-web/internal/feature/marketing"
)

const MOD_FEATURE string = "FEATURE"

type Module struct {
	config   *module.Config
	registry *module.Registry
}

func New(c *module.Config, r *module.Registry) *Module {
	return &Module{
		registry: r,
	}
}

func (m *Module) Name() string {
	return MOD_FEATURE
}

func (m *Module) Initialize() error {
	m.registry.Register(auth.New(m.config, m.registry))
	m.registry.Register(marketing.New(m.config, m.registry))
	return nil
}
