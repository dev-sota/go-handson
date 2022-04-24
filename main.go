package main

import (
	"log"
	"net/http"

	"github.com/dev-sota/go-handson/controller"
	"github.com/dev-sota/go-handson/database"
	"github.com/go-chi/chi"
)

func newRouter() http.Handler {
	r := chi.NewRouter()
	u := controller.NewUser()

	r.Get("/users/{id}", u.Get)
	return r
}

func main() {
	err := database.Init()
	if err != nil {
		log.Fatalln(err)
	}

	h := newRouter()
	http.ListenAndServe(":3000", h)
}
