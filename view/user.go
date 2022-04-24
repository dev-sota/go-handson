package view

import "github.com/dev-sota/go-handson/model"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(m *model.User) *User {
	return &User{
		Name:  m.Name,
		Email: m.Email,
	}
}
