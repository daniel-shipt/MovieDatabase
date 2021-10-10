package repo

import (
	"MovieDatabase/entities"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//type MovieRepo interface{
//	AddMovie(m entities.Movie) error
//	ViewAll() (DataBase, error)
//	FindById(id string) (entities.Movie, error)
//	DeleteMovie(id string) error
//	UpdateMovie(id string, mv entities.Movie) error
//}

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
	return entities.Movie{}, errors.New("movie not found")
}

func (r Repo) DeleteMovie(id string) error {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	movies := DataBase{}
	newDb := DataBase{}
	err = json.Unmarshal(file, &movies)

	dbSize := len(movies.Movies)
	for _, val := range movies.Movies {
		if val.Id == id {
			continue
		} else {
			newDb.Movies = append(newDb.Movies, val)
		}
	}

	if len(newDb.Movies) == dbSize {
		return errors.New("failed to delete movie - does not exist")
	}

	movieBytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) UpdateMovie(id string, mv entities.Movie) error {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	movies := DataBase{}
	err = json.Unmarshal(file, &movies)
	if err != nil {
		return err
	}

	for i, val := range movies.Movies {
		if val.Id == id {
			movies.Movies[i] = mv
		}
	}

	movieBytes, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
