package main

import (
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
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	err := http.ListenAndServe(":8888", r)
	log.Fatalf("Error %e start server", err)
}
