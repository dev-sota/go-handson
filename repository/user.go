package repository

import (
	"github.com/dev-sota/go-handson/model"
)

type User interface {
	FindByEmail(email string) (*model.User, error)
	Find(id int64) (*model.User, error)
	Create(name, email string, age int) (*model.User, error)
}
