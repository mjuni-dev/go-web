package gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	ctx *gin.Context
}

func (c *Context) Ctx() context.Context {
	return c.ctx
}

func (c *Context) Request() *http.Request {
	return c.ctx.Request
}

func (c *Context) Response() *http.Response {
	return c.ctx.Request.Response
}

func (c *Context) ResponseWriter() *http.ResponseWriter {
	return nil
}

func (c *Context) String(code int, s string) error {
	c.ctx.String(code, s)
	return nil
}

func (c *Context) HTML(code int, html string) error {
	c.ctx.Header("Content-Type", "text/html")
	c.ctx.String(code, html)
	return nil
}
