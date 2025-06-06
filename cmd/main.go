package main

import (
	"log"

	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server/factory"
	"github.com/mjuni.dev/go-web/internal/core/server/server"
	"github.com/mjuni.dev/go-web/internal/feature"
	"github.com/mjuni.dev/go-web/web"
)

func main() {
	f := factory.NewServerFactory()
	s := f.EchoServer()
	// s := f.GinServer()
	router := s.Router()

	registry := module.NewRegistry()
	initializeFeatureModules(registry)
	initializeWebModules(registry, router)

	log.Fatal(s.Start(":8080"))
}

func initializeFeatureModules(r *module.Registry) {
	c := module.Config{
		ConnectionString: "",
	}
	r.Register(feature.New(&c, r))
}

func initializeWebModules(registry *module.Registry, router server.Router) {
	registry.Register(web.New(registry, router))
}
