package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dev-sota/go-handson/usecase"
	"github.com/dev-sota/go-handson/view"
	"gorm.io/gorm"
)

type user struct {
	uu usecase.User
}

func NewUser(db *gorm.DB) *user {
	return &user{
		uu: *usecase.NewUser(db),
	}
}

func (u user) Signup(w http.ResponseWriter, r *http.Request) {
	name := "dev-sota"
	email := "test@exam.com"
	age := 25

	m, _ := u.uu.Signup(name, email, age)

	v := view.NewUser(m)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
