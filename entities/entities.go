package entities

type Movie struct {
	Id 			string
	Title 		string
	Genre 		[]string
	Description string
	Director 	string
	Actors 		[]string
	Rating 		float64
}

func(m *Movie) UpdateMovie(id string, mov Movie){
	switch {
	case m.Id != id:
		m.Id = id
	case m.Title != mov.Title:
		m.Title = mov.Title
	case m.Description != mov.Description:
		m.Description = mov.Description
	case m.Director != mov.Director:
		m.Director = mov.Director
	case m.Rating != mov.Rating:
		m.Rating = mov.Rating
	}

	m.Genre = mov.Genre
	m.Actors = mov.Actors
}