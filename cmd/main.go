package main

import (
	"log"

	"github.com/mjuni.dev/go-web/internal/core/module"
	"github.com/mjuni.dev/go-web/internal/core/server"
	"github.com/mjuni.dev/go-web/internal/features/auth"
	"github.com/mjuni.dev/go-web/internal/features/marketing"
)

func main() {
	f := server.NewServerFactory()
	s := f.EchoServer()
	// s := f.GinServer()
	// r := s.Router()

	initializeModules()

	log.Fatal(s.Start(":8080"))
}

func initializeModules() {
	c := module.Config{
		ConnectionString: "",
	}
	r := module.NewRegistry()

	r.Register(auth.New(&c))
	r.Register(marketing.New(&c))
}
