package module

import "fmt"

type Registry struct{}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(m Module) error {
	fmt.Printf(" >> Initializing %s...\n", m.Name())
	return m.Initialize()
}
