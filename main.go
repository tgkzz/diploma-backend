package main

import (
	"diploma/internal/server"
	"log"
	"net/http"
)

func main() {
	routes := server.MakeRoutes()

	log.Fatal((http.ListenAndServe(":4000", routes)))
}
