package pricing

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/features/marketing/pricing/templates"
)

type Module struct {
	router server.Router
}

func New(r server.Router) *Module {
	return &Module{
		router: r,
	}
}

func (m *Module) Initialize() error {
	fmt.Println(" >> >> Initializing WEB-PRICING sub-module...")
	m.router.GET("/pricing", m.handleGetPricing)
	return nil
}

func (m *Module) handleGetPricing(c server.Context) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	component := templates.PricingTiers()
	component.Render(c.Ctx(), buf)
	return c.HTML(200, buf.String())
}
