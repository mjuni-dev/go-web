package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

type Router struct {
	echo *echo.Echo
}

func (r *Router) GET(path string, handler server.HandlerFunc, middlewares ...server.MiddlewareFunc) {
	r.echo.GET(path, func(c echo.Context) error {
		return handler(&Context{ctx: c})
	}, r.convertMiddlewares(middlewares...)...)
}

func (r *Router) convertMiddlewares(middlewares ...server.MiddlewareFunc) []echo.MiddlewareFunc {
	m := make([]echo.MiddlewareFunc, len(middlewares))

	return m
}
