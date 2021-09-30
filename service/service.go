package service

import (
	"MovieDatabase/entities"
	"MovieDatabase/repo"
	"github.com/google/uuid"
)

type Service struct {
	Repo repo.Repo
}

func DoService(r repo.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) AddMovie(m entities.Movie) error {
	m.Id = uuid.New().String()

	err := s.Repo.AddMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ViewAll() (repo.DataBase, error) {
	db, err := s.Repo.ViewAll()
	if err != nil {
		return db, err
	}
	return db, nil
}

func (s Service) FindById(id string) (entities.Movie, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}
