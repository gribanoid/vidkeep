package main

import (
	"log"

	"github.com/gribanoid/vidkeep/config"
	"github.com/gribanoid/vidkeep/internal/server"
)

var c config.Config

func init() {
	err := config.ReadFile(&c)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	srv := server.NewHTTPServer(&c)
	log.Printf("Server started on http://localhost:%s", c.Server.Port)
	log.Fatal(srv.ListenAndServe())
}
