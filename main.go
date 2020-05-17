// Package main provides ...
package main

import (
	"log"
	"net/http"

	"github.com/a0a0o0/colly-server/api"
)

func main() {
	server := &http.Server{
		Addr:    ":80",
		Handler: &api.ServerHandler{},
	}
	log.Fatal(server.ListenAndServe())
}
