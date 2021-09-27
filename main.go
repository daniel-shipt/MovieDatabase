package main

import (
	"MovieDatabase/handlers"
	"log"
)

func main() {
	server := handlers.NewServer()

	log.Fatal(server.ListenAndServe())
}