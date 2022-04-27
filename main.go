package main

import (
	"net/http"

	"github.com/dev-sota/go-handson/controller"
	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:password@tcp(127.0.0.1:3306)/go-handson?charset=utf8mb4&parseTime=True&loc=Local"

func newRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	u := controller.NewUser(db)

	r.Post("/users/signup", u.Signup)
	return r
}

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	h := newRouter(db)
	http.ListenAndServe(":3000", h)
}
