package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mjuni.dev/go-web/internal/core/server/interfaces"
)

type Router struct {
	echo *echo.Echo
}

func (r *Router) GET(path string, handler interfaces.HandlerFunc, middlewares ...interfaces.MiddlewareFunc) {
	r.echo.GET(path, func(c echo.Context) error {
		return handler(&Context{ctx: c})
	}, r.convertMiddlewares(middlewares...)...)
}

func (r *Router) convertMiddlewares(middlewares ...interfaces.MiddlewareFunc) []echo.MiddlewareFunc {
	m := make([]echo.MiddlewareFunc, len(middlewares))

	return m
}
