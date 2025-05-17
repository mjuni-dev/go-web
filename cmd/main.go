package main

import (
	"log"

	"github.com/mjuni.dev/go-web/internal/core/server"
	"github.com/mjuni.dev/go-web/internal/features/auth"
)

func main() {
	f := server.New()
	s := f.EchoServer()
	// s := f.GinServer()
	// r := s.Router()

	initializeModules()
	log.Fatal(s.Start(":3000"))
}

func initializeModules() {
	auth.New().Initialize()
}
