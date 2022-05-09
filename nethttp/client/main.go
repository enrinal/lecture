package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeHandler struct {
	Timezone string
	Format   string
}

func (h TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation(h.Timezone)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	t := time.Now().In(loc).Format(h.Format)
	w.Write([]byte(fmt.Sprintf("Today is %v in Timezone %s", t, h.Timezone)))
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("First Middleware")
		next.ServeHTTP(w, r)
		fmt.Println("First Middleware End")
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing secondMiddleware")
		if r.Method != http.MethodGet {
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("Executing secondMiddleware again")
	})
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	handler := http.HandlerFunc(helloWorldHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: firstMiddleware(secondMiddleware(handler)),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Println("Server started on port 8080")
}
