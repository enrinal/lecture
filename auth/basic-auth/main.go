package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	USERNAME = "aditira"
	PASSWORD = "aditira123"
)

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func init() {
	students = append(students, &Student{Id: "001", Name: "Gina", Grade: 2})
	students = append(students, &Student{Id: "002", Name: "Indri", Grade: 2})
	students = append(students, &Student{Id: "003", Name: "Dwi", Grade: 3})
}

func main() {
	http.HandleFunc("/student", ActionStudent) // router creation

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server started at http://localhost:9000")
	server.ListenAndServe()
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	log.Println("Authorization : ", r.Header.Get("Authorization"))
	// check basic auth
	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized.\n"))
		return false
	}
	// check username and password
	if username != USERNAME || password != PASSWORD {
		w.WriteHeader(401)
		w.Write([]byte("wrong username/password.\n"))
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !AllowOnlyGET(w, r) {
		return
	}

	if !Auth(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, s := range students {
		if s.Id == id {
			return s
		}
	}

	return nil
}
