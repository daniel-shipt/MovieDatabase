package service

import (
	"MovieDatabase/entities"
	"MovieDatabase/repo"
	"errors"
	"github.com/google/uuid"
)

//type MovieService interface {
//	AddMovie(m entities.Movie) error
//	ViewAll() (repo.DataBase, error)
//	FindById(id string) (entities.Movie, error)
//	DeleteMovie(id string) error
//	UpdateMovie(id string, m entities.Movie) error
//}

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

	if m.Rating >= 0 && m.Rating <= 10 {
		err := s.Repo.AddMovie(m)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid rating")
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

func (s Service) DeleteMovie(id string) error {
	err := s.Repo.DeleteMovie(id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateMovie(id string, m entities.Movie) error {
	err := s.Repo.UpdateMovie(id, m)
	if err != nil {
		return err
	}
	return nil
}
