package handlers

import (
	"MovieDatabase/entities"
	"MovieDatabase/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil{
		fmt.Println(err)
	}

	mv.GetId()

	mvdb := repo.DataBase{}
	mvdb.AddToDb(mv)

	movieBytes, err := json.MarshalIndent(mvdb, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", movieBytes, 0644)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(movieBytes)
}