package pricing

const FEATURE_MARKETING_PRICING string = "FEATURE.MARKETING"

type Module struct {
}

func New() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return FEATURE_MARKETING_PRICING
}

func (m *Module) Initialize() error {
	return nil
}
