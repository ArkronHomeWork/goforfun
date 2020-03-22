package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

type postData struct {
	userName     string
	userPassword string
}

func (data *postData) toStruct(rawData []byte) error {
	return json.Unmarshal(rawData, &data)
}

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
		data := new(postData)
		rawData, err := ioutil.ReadAll(r.Body)
		err = data.toStruct(rawData)
		if err != nil {
			log.Print(err)
			panic(err)
		}
	})
}
