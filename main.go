package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

/*func main() {
	err := http.ListenAndServe(":8888", nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Fatalf("Error %e start server", err)
}*/

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
}
