package gin

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Context struct {
	ctx *gin.Context
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) String(code int, s string) error {
	c.ctx.String(code, s)
	return nil
}

func (c *Context) Value(key any) any {
	return nil
}

func (c *Context) Deadline() (t time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}
