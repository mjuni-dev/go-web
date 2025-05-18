package echo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Context struct {
	ctx echo.Context
}

func (c *Context) Ctx() context.Context {
	return c.ctx.Request().Context()
}

func (c *Context) Request() *http.Request {
	return c.ctx.Request()
}

func (c *Context) Response() *http.Response {
	return c.ctx.Request().Response
}

func (c *Context) ResponseWriter() *http.ResponseWriter {
	return &c.ctx.Response().Writer
}

func (c *Context) String(code int, s string) error {
	return c.ctx.String(code, s)
}

func (c *Context) HTML(code int, html string) error {
	return c.ctx.HTML(code, html)
}
