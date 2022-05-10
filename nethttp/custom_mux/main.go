package main

import "net/http"

type StudentHandler struct{}

func (h *StudentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World From StudentHandler"))
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World From TeacherHandler"))
}

func main() {
	mux := http.NewServeMux()

	// custom handler
	sHandler := &StudentHandler{}
	mux.Handle("/student", sHandler)

	// default handler
	mux.HandleFunc("/teacher", TeacherHandler)

	http.ListenAndServe(":8080", mux)
}
