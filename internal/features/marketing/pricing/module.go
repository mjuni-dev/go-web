package pricing

import (
	"fmt"
)

type Module struct {
}

func New() *Module {
	return &Module{}
}

func (m *Module) Initialize() error {
	fmt.Println(" >> >> Initializing PRICING sub-feature...")
	return nil
}
