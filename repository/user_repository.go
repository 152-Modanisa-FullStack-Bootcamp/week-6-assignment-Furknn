package repository

import (
	"errors"
	"week-6-assignment-Furknn/config"
	"week-6-assignment-Furknn/model"
)

var userStorage = []model.User{
	{
		Username: "Furkan",
		Balance:  1,
	},
}

type IUserRepository interface {
	FindAllUsers() (model.Users, error)
	FindUser(username string) (model.User, error)
	AddUser(username string) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (p *UserRepository) FindAllUsers() (model.Users, error) {
	return userStorage, nil
}

func (p *UserRepository) FindUser(username string) (model.User, error) {
	var user model.User
	for _, usr := range userStorage {
		if usr.Username == username {
			user = usr
			return user, nil
		}
	}
	return user, errors.New("user not found")
}

func (p *UserRepository) AddUser(username string) (model.User, error) {
	userStorage := append(userStorage, model.User{Username: username, Balance: config.C.InitialBalanceAmount})
	return userStorage[len(userStorage)-1], nil
}

func (p *UserRepository) UpdateUser(user model.User) (model.User, error) {
	var index int
	found := false
	for i, usr := range userStorage {
		if usr.Username == user.Username {
			index = i
			found = true
		}
	}
	if found {
		userStorage[index] = user
		return user, nil
	}
	return user, errors.New("user not found")

}
