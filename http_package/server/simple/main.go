package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whatsupp!\n")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
