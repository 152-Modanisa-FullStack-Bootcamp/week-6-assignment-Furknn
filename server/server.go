package server

import (
	"fmt"
	"net/http"
	"week-6-assignment-Furknn/handler"
	"week-6-assignment-Furknn/repository"
	"week-6-assignment-Furknn/service"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/", userHandler.User)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
