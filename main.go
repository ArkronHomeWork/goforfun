package main

import (
	"github.com/ArkronHomeWork/goforfun/srv"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	service := srv.GetUserService()
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Get("/", service.GetAllUsers)
		r.Post("/", service.SaveNewUser)

		r.Get("/{id}", service.GetUserById)
	})
	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatalf("Error %e start server", err)
	}
}
