package marketing

import (
	"fmt"

	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/features/marketing/pricing"
)

type Module struct {
	config   *module.Config
	registry *module.Registry
	// repository Repository
	// service    *Service
	// handler    *Handler
}

func New(c *module.Config, r *module.Registry) *Module {
	return &Module{
		config:   c,
		registry: r,
	}
}

func (m *Module) Initialize() error {
	fmt.Println(" >> Initializing MARKETING feature...")
	m.registry.Register(pricing.New())
	// m.repository = NewRepository()
	// m.service = NewService(m.repository)
	// m.handler = NewHandler(m.service)

	return nil
}
