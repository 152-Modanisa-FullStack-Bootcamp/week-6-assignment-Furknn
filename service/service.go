package service

import (
	"encoding/json"
	"net/http"
	"week-6-assignment-Furknn/config"
	"week-6-assignment-Furknn/model"
	"week-6-assignment-Furknn/repository"
)

type IUserService interface {
	GetUsers() (model.Users, error)
	GetUser(r *http.Request) (model.User, error)
	PostUser(r *http.Request) (model.User, error)
	PutUser(r *http.Request) (model.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (p *UserService) GetUser(r *http.Request) (model.User, error) {
	username := r.URL.Query().Get("username")
	user, err := p.repo.FindUser(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *UserService) PutUser(r *http.Request) (model.User, error) {
	username := r.URL.Query().Get("username")
	var user model.User
	findUser, err := p.repo.FindUser(username)
	if err != nil {
		addUser, err := p.repo.AddUser(username)
		if err != nil {
			return model.User{}, err
		}
		user = addUser
		return user, nil
	}
	user = findUser
	return user, nil
}

func (p *UserService) PostUser(r *http.Request) (model.User, error) {
	username := r.URL.Query().Get("username")
	decoder := json.NewDecoder(r.Body)
	var balance model.Balance

	err := decoder.Decode(&balance)
	if err != nil {
		return model.User{}, err
	}
	value := balance.Balance

	user, err := p.repo.FindUser(username)
	if err != nil {
		return user, err
	}

	// add value
	user.Balance = user.Balance + value

	// check if balance is smaller than minimum
	if user.Balance < config.C.MinumumBalanceAmount {
		user.Balance = config.C.MinumumBalanceAmount
	}

	updateUser, err := p.repo.UpdateUser(user)
	if err != nil {
		return user, err
	}
	return updateUser, nil

}
func (p *UserService) GetUsers() (model.Users, error) {
	return p.repo.FindAllUsers()
}
