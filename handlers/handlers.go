package handlers

import (
	"MovieDatabase/entities"
	"MovieDatabase/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MovieHandler struct {
	Serv service.Service
}

func NewMovieHandler(s service.Service) MovieHandler {
	return MovieHandler{
		Serv: s,
	}
}

func (mov MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = mov.Serv.AddMovie(mv)
	if err != nil {
		switch err.Error() {
		case "movie already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (mov MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movDb, err := mov.Serv.ViewAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	movieDb, err := json.MarshalIndent(movDb, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(movieDb)
}

func (mov MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	movById, err := mov.Serv.FindById(id)
	if err != nil {
		switch err.Error() {
		case "movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	movie, err := json.MarshalIndent(movById, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(movie)
}

func (mov MovieHandler) DeleteMov(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mov.Serv.DeleteMovie(id)
	if err != nil {
		switch err.Error() {
		case "failed to delete movie - does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
