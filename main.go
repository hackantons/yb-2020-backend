package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":5080", nil))
}

// Just for testing purposes
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
