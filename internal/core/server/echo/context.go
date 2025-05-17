package echo

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	ctx echo.Context
}

// func (c *Context) Err() error {
// 	return nil
// }

func (c *Context) String(code int, s string) error {
	return c.ctx.String(code, s)
}

// func (c *Context) Value(key any) any {
// 	return nil
// }

// func (c *Context) Deadline() (t time.Time, ok bool) {
// 	return
// }

// func (c *Context) Done() <-chan struct{} {
// 	return nil
// }
