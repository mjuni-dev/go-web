package register

import (
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/feature/auth/register/templates/pages"
	"github.com/mjuni.dev/go-web/web/utils"
)

const MOD_WEB_MARKETING string = "WEB.AUTH.REGISTER"

type Module struct {
	router server.Router
}

func New(r server.Router) *Module {
	return &Module{
		router: r,
	}
}

func (m *Module) Initialize() error {
	m.router.GET("/auth/register", m.handleGetAbout)
	return nil
}

func (m *Module) Name() string {
	return MOD_WEB_MARKETING
}

func (m *Module) handleGetAbout(c server.Context) error {
	return utils.Render(c, 200, pages.Register())
}
