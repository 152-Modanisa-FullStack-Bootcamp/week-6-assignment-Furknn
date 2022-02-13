package service

import (
	"week-6-assignment-Furknn/model"
	"week-6-assignment-Furknn/repository"
)

type IUserService interface {
	Users() (model.Users, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (p *UserService) Users() (model.Users, error) {
	return p.repo.FindAllUsers()
}
