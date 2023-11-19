package service

import (
	"ms-go-initial/db"
)

type IService interface {
	CreateUser(userName string, status bool) error
	UpdateUser(userId uint, userName string, status bool) error
}

type service struct {
	repo db.IRepository
}

func NewService(repo db.IRepository) IService {
	return &service{repo: repo}
}

func (s service) CreateUser(userName string, status bool) error {

	err := s.repo.SaveUser(userName, status)

	if err != nil {
		return err
	}

	return nil

}

func (s service) UpdateUser(userId uint, userName string, status bool) error {

	err := s.repo.UpdateUser(userId, userName, status)

	if err != nil {
		return err
	}

	return nil

}
