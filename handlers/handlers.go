package handlers

import (
	"MovieDatabase/entities"
	"MovieDatabase/repo"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}
	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	db, err := repo.AddMovie(mv)
	if err != nil {
		fmt.Println(err)
	}

	movieBytes, err := json.MarshalIndent(db, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(movieBytes)
}
