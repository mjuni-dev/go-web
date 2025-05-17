package module

type Registry struct{}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(m Module) error {
	return m.Initialize()
}
