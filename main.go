package main

import (
	"fmt"
	"github.com/ArkronHomeWork/goforfun/repository"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	_, err := databaseConnect()
	if err != nil {
		log.Fatalf("Error %e database error", err)

	}
	r := chi.NewRouter()
	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	err = http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatalf("Error %e start server", err)
	}

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		data := new(repository.PostData)
		rawData, err := ioutil.ReadAll(r.Body)
		err = data.ToStruct(rawData)
		if err != nil {
			log.Print(err)
			panic(err)
		}
	})
}
