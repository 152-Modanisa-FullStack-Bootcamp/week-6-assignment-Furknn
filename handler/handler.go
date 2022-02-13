package handler

import (
	"encoding/json"
	"net/http"
	"week-6-assignment-Furknn/service"
)

type IHandler interface {
	GetUsers(writer http.ResponseWriter, request *http.Request)
}

type Handler struct {
	service service.IUserService
}

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	products, err := h.service.Users()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productsBytes, _ := json.Marshal(products)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(productsBytes)
}

func NewUserHandler(service service.IUserService) IHandler {
	return &Handler{service: service}
}
