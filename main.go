package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sweetcheetah/mqtt-api/server"
)

func main() {
	router := mux.NewRouter()

	rest := server.NewServer(router)

	log.Printf("Starting mqtt-api...")

	log.Fatal(rest.Start())
}
