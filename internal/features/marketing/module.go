package marketing

import (
	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/features/marketing/pricing"
)

const FEATURE_MARKETING string = "FEATURE.MARKETING"

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

func (m *Module) Name() string {
	return FEATURE_MARKETING
}

func (m *Module) Initialize() error {
	m.registry.Register(pricing.New())
	// m.repository = NewRepository()
	// m.service = NewService(m.repository)
	// m.handler = NewHandler(m.service)

	return nil
}
