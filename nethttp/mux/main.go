package main

import (
	"log"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func middleWareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer 12345" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func middleWareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		log.Println(r.Method)
		log.Println(r.Header)
		log.Println("The request is done")

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.DefaultServeMux
	// set routes for mux
	mux.HandleFunc("/simple", simpleHandler)

	var handler http.Handler = mux

	// add middleware
	handler = middleWareLog(handler)
	handler = middleWareAuth(handler)

	http.ListenAndServe(":8080", handler)
}
