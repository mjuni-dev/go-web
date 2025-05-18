package utils

import (
	"github.com/a-h/templ"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

func Render(c server.Context, statuscode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(c.Ctx(), buf); err != nil {
		return err
	}

	return c.HTML(statuscode, buf.String())
}
