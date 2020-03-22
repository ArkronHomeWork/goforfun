package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	err := http.ListenAndServe(":8888", nil)
	log.Fatalf("Error %e start server", err)
}
