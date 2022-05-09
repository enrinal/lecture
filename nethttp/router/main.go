package main

import (
	"log"
	"net/http"
)

/*
	http://localhost:8080/ --> handlerIndex
	http://localhost:8080/index --> handlerIndex
*/

func main() {
	handlerIdex := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World From Router 1"}`))
	}

	handlerIndex1 := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World From Router 2"}`))
	}

	http.HandleFunc("/", handlerIdex)
	http.HandleFunc("/index", handlerIdex)
	http.HandleFunc("/index1", handlerIndex1)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World From Router 3"}`))
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
