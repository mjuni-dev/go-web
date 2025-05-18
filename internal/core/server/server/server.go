package server

import (
	"context"
	"net/http"
)

type Server interface {
	Start(address string) error
	Router() Router
}

type Router interface {
	GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc)
}

type Context interface {
	Ctx() context.Context
	Request() *http.Request
	Response() *http.Response
	ResponseWriter() *http.ResponseWriter
	String(code int, s string) error
	HTML(code int, html string) error
}

type HandlerFunc func(ctx Context) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
