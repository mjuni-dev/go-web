package auth

import "fmt"

type Module struct{}

func New() *Module {
	return &Module{}
}

func (m *Module) Initialize() error {
	fmt.Println("Initializing auth feature...")
	return nil
}
