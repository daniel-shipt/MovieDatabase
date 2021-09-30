package repo

import (
	"MovieDatabase/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DataBase struct {
	Movies []entities.Movie
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
	movieStruct := DataBase{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &movieStruct)
	if err != nil {
		return err
	}

	movieStruct.Movies = append(movieStruct.Movies, m)

	movieBytes, err := json.MarshalIndent(movieStruct, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) ViewAll() (DataBase, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	db := DataBase{}
	err = json.Unmarshal(file, &db)
	if err != nil {
		return db, err
	}

	return db, nil
}

func (r Repo) FindById(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	movies := DataBase{}
	err = json.Unmarshal(file, &movies)

	match := entities.Movie{}

	for _, val := range movies.Movies {
		if val.Id == id {
			match = val
			return match, nil
		}
	}
	return match, nil
}
