package model

import (
	"github.com/dev-sota/go-handson/database"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func (u *User) FindByEmail(email string) error {
	return database.Get().
		Where("email = ?", email).
		Find(u).
		Error
}

func (u *User) Find(id int64) error {
	return database.Get().First(u, id).Error
}

func (u *User) Create() error {
	return database.Get().Create(u).Error
}

func (u *User) Update() error {
	return database.Get().Save(u).Error
}

func (u *User) Delete() error {
	return database.Get().Delete(u).Error
}
