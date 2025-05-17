package echo

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	ctx echo.Context
}

func (c *Context) String(code int, s string) error {
	return c.ctx.String(code, s)
}
