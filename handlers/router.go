package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.GetById).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.DeleteMov).Methods("DELETE")

	return r
}
