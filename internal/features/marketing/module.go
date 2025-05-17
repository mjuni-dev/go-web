package marketing

import (
	"fmt"

	"github.com/mjuni.dev/go-web/internal/core/module"
)

type Module struct {
	config *module.Config
}

func New(c *module.Config) *Module {
	return &Module{
		config: c,
	}
}

func (m *Module) Initialize() error {
	fmt.Println("Initializing marketing feature...")
	return nil
}
