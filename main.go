package main

import (
	"MovieDatabase/handlers"
	"MovieDatabase/repo"
	"MovieDatabase/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	file := "moviedb.json"

	ext := filepath.Ext(file)
	if ext != ".json" {
		log.Fatal("Invalid File Extension")
	}

	repository := repo.NewRepo(file)
	serv := service.DoService(repository)
	handler := handlers.NewMovieHandler(serv)
	router := handlers.ConfigRouter(handler)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(server.ListenAndServe())
}
