package auth

import (
	"fmt"

	"github.com/mjuni.dev/go-web/internal/core/module"
)

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

func (m *Module) Initialize() error {
	fmt.Println(" >> Initializing AUTH feature...")
	return nil
}
