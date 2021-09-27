package repo

import (
	"MovieDatabase/entities"
)

type DataBase struct{
	Movies []entities.Movie `json:"Movies"`
}

func (db *DataBase) AddToDb(movie entities.Movie) {
	db.Movies = append(db.Movies, movie)
}

