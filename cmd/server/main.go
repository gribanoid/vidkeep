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
	srv := server.NewHTTPServer(":"+c.Server.Port, c.Storage.Path)
	log.Fatal(srv.ListenAndServe())
}
