package repo

import (
	"MovieDatabase/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DataBase struct {
	Movies []entities.Movie `json:"Movies"`
}

func AddMovie(m entities.Movie) (DataBase, error) {
	file, err := ioutil.ReadFile("moviedb.json")
	if err != nil {
		fmt.Println(err)
	}
	m.GetId()
	dbSlice := DataBase{}
	err = json.Unmarshal(file, &dbSlice)
	if err != nil {
		fmt.Println(err)
	}
	dbSlice.Movies = append(dbSlice.Movies, m)

	movieBytes, err := json.MarshalIndent(dbSlice, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", movieBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return dbSlice, nil
}
