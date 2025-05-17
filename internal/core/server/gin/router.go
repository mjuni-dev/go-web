package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mjuni.dev/go-web/internal/core/server/interfaces"
)

type Router struct {
	gin *gin.Engine
}

func (r *Router) GET(path string, handler interfaces.HandlerFunc, middlewares ...interfaces.MiddlewareFunc) {
	ginHandler := func(c *gin.Context) {
		err := handler(&Context{ctx: c})
		if err != nil {
		}
	}
	ginMiddlewares := r.convertMiddlewares(middlewares...)
	r.gin.GET(path, append(ginMiddlewares, ginHandler)...)
	// r.gin.GET(path, func(c *gin.Context) {
	// 	handler(&Context{ctx: c})
	// })
}

func (r *Router) convertMiddlewares(middlewares ...interfaces.MiddlewareFunc) []gin.HandlerFunc {
	ginMiddlewares := make([]gin.HandlerFunc, len(middlewares))

	return ginMiddlewares
}
