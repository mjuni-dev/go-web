package interfaces

type Server interface {
	Start(address string) error
	Router() Router
}

type Router interface {
	GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc)
}

type Context interface {
	// Err() error
	String(code int, s string) error

	// Value(key any) any

	// Deadline() (t time.Time, ok bool)
	// Done() <-chan struct{}
}

type HandlerFunc func(ctx Context) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
