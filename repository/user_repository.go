package repository

import "week-6-assignment-Furknn/model"

type IUserRepository interface {
	FindAllUsers() (model.Users, error)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (p *UserRepository) FindAllUsers() (model.Users, error) {
	Users := model.Users{}
	Users = append(Users, model.User{
		Username: "Furkan",
		Balance:  1,
	})

	return Users, nil
}
