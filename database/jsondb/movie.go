package jsondb

import (
	"MovieDatabase/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type DataBase struct {
	Movies []entities.Movie
	//MovieMap map[string]entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo {
	return Repo{
		Filename: f,
	}
}

func (r Repo) AddMovie(m entities.Movie) error {
	movieSlice := DataBase{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &movieSlice)
	if err != nil {
		return err
	}

	for _, val := range movieSlice.Movies {
		if val.Title == m.Title {
			return errors.New("movie already exists")
		}
	}

	movieSlice.Movies = append(movieSlice.Movies, m)

	movieBytes, err := json.MarshalIndent(movieSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}


