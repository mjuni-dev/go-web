package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
)

type Router struct {
	gin *gin.Engine
}

func (r *Router) GET(path string, handler server.HandlerFunc, middlewares ...server.MiddlewareFunc) {
	r.gin.GET(path, func(c *gin.Context) {
		handler(&Context{ctx: c})
	})
}

func (r *Router) convertMiddlewares(middlewares ...server.MiddlewareFunc) []gin.HandlerFunc {
	ginMiddlewares := make([]gin.HandlerFunc, len(middlewares))

	return ginMiddlewares
}
