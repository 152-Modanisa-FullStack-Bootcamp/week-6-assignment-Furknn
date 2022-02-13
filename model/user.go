package model

type User struct {
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

type Users []User
