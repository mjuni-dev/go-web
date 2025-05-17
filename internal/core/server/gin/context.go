package gin

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	ctx *gin.Context
}

func (c *Context) String(code int, s string) error {
	c.ctx.String(code, s)
	return nil
}
