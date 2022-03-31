package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	Admin     bool   `json:"admin"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		user := User{
			FirstName: "John Rogu Rogu",
			Admin:     false,
		}
		jsonResp, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResp)
		return
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "POST request to the homepage")
		return
	} else if r.Method == "PUT" {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "PUT request to the homepage")
		return
	} else if r.Method == "DELETE" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "DELETE request to the homepage")
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
