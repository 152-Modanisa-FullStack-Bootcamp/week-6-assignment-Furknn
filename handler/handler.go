package handler

import (
	"encoding/json"
	"net/http"
	"week-6-assignment-Furknn/service"
)

type IHandler interface {
	User(writer http.ResponseWriter, request *http.Request)
}

type Handler struct {
	service service.IUserService
}

func (h Handler) User(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		username := r.URL.Query().Get("username")

		// Get Users
		if username == "" {
			users, err := h.service.GetUsers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			userBytes, _ := json.Marshal(users)
			w.Header().Add("content-type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(userBytes)

			// Get User
		} else {
			user, err := h.service.GetUser(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			userBytes, _ := json.Marshal(user)
			w.Header().Add("content-type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(userBytes)
		}

		// Put User
	} else if r.Method == http.MethodPut {
		user, err := h.service.PutUser(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userBytes, _ := json.Marshal(user)
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(userBytes)

		// Post User
	} else if r.Method == http.MethodPost {
		user, err := h.service.PostUser(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		userBytes, _ := json.Marshal(user)
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(userBytes)
	}
}

func NewUserHandler(service service.IUserService) IHandler {
	return &Handler{service: service}
}
