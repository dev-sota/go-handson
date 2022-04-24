package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-sota/go-handson/model"
	"github.com/dev-sota/go-handson/view"
	"github.com/go-chi/chi"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

func (u user) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return
	}

	var m *model.User
	if err := m.Find(int64(id)); err != nil {
		return
	}

	v := view.NewUser(m)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func (u user) Signup(w http.ResponseWriter, r *http.Request) {
	name := "dev-sota"
	email := "test@exam.com"

	var m *model.User
	if err := m.FindByEmail(email); err != nil {
		return
	}
	if m.ID != 0 {
		fmt.Println("already registerd")
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	// ビジネスロジック
	// if name == "ito" {
	// 	fmt.Println("クーポン発行")
	// }
	// if age < 20 {
	// 	fmt.Println("クーポン発行")
	// }
	// if birthDayMonth == time.Now().Month() {
	// 	fmt.Println("クーポン発行")
	// }

	usr := model.NewUser(name, email)
	usr.Create()

	v := view.NewUser(usr)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
