package interfaces

type Server interface {
	Start(address string) error
	Router() Router
}

type Router interface {
	GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc)
}

type Context interface {
	String(code int, s string) error
}

type HandlerFunc func(ctx Context) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
