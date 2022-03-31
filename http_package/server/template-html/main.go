package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var testTemplate *template.Template

/*
if User != nil && User.Admin != false {
	Welcome .User.FirstName
}else {
	Access Denied!
*/

const htmlTemplate = `
{{if and .User .User.Admin}}
  <h1>  Welcome {{ .User.FirstName }}, </h1>
   <h2> You are an admin user! </h2>
{{else}}
    <h1> Access denied! </h1>
{{end}}
`

type User struct {
	FirstName string
	Admin     bool
}

type ViewData struct {
	User User
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{
		User: User{
			FirstName: "John Rogu",
			Admin:     true,
		},
	}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error

	testTemplate, err = template.New("htmlTemplate").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/user", handlerUser)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}
