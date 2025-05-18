package auth

const MOD_WEB_AUTH string = "WEB.AUTH"

type Module struct{}

func New() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return MOD_WEB_AUTH
}

func (m *Module) Initialize() error {
	return nil
}
