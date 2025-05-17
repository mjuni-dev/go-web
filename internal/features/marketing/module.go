package marketing

import (
	"fmt"

	"github.com/mjuni.dev/go-web/internal/core/module"
)

type Module struct {
	config     *module.Config
	repository Repository
	service    *Service
	handler    *Handler
}

func New(c *module.Config) *Module {
	return &Module{
		config: c,
	}
}

func (m *Module) Initialize() error {
	fmt.Println("Initializing marketing feature...")
	m.repository = NewRepository()
	m.service = NewService(m.repository)
	m.handler = NewHandler(m.service)

	return nil
}
