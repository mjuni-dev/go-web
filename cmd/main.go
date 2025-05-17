package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjuni.dev/go-web/internal/core/server"
	"github.com/mjuni.dev/go-web/internal/core/server/interfaces"
)

func main() {
	f := server.New()
	s := f.EchoServer()
	// s := f.GinServer()
	r := s.Router()
	r.GET("/", func(c interfaces.Context) error {
		fmt.Printf(" >> GET\n")
		return c.String(http.StatusOK, "Hello, world!\n")
	})
	log.Fatal(s.Start(":3000"))
}
