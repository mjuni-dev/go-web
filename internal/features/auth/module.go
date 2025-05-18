package auth

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
)

const FEATURE_AUTH string = "FEATURE.AUTH"

type Module struct {
	config   *module.Config
	registry *module.Registry
}

func New(c *module.Config, r *module.Registry) *Module {
	return &Module{
		config:   c,
		registry: r,
	}
}
func (m *Module) Name() string {
	return FEATURE_AUTH
}

func (m *Module) Initialize() error {
	return nil
}
