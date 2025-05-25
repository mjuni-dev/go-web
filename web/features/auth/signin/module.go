package signin

import (
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/web/features/auth/signin/templates/pages"
	"github.com/mjuni.dev/go-web/web/utils"
)

const MOD_WEB_MARKETING string = "WEB.AUTH.SIGNIN"

type Module struct {
	router server.Router
}

func New(r server.Router) *Module {
	return &Module{
		router: r,
	}
}

func (m *Module) Initialize() error {
	m.router.GET("/auth/signin", m.handleGetAbout)
	return nil
}

func (m *Module) Name() string {
	return MOD_WEB_MARKETING
}

func (m *Module) handleGetAbout(c server.Context) error {
	return utils.Render(c, 200, pages.SignIn())
}
