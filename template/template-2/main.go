package main

import (
	"fmt"
	"os"
	"text/template"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	s := Student{
		Name:  "Budi",
		Age:   18,
		Score: 100,
	}

	t, err := template.New("student").Parse(`{{ if .Name}}Halo {{.Name}}! {{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, s)
	fmt.Println("")
	t2, err := template.New("student1").Parse(`{{if (gt .Age 18)}}Halo {{.Name}}! {{else}}Halo {{.Name}}! {{end}}`)
	if err != nil {
		panic(err)
	}
	err = t2.Execute(os.Stdout, s)
	fmt.Println("")

	t3, err := template.New("student2").Parse(`{{if (gt .Score 90.00)}}Halo {{.Name}} Anda Lulus! {{else}} Halo {{.Name}} Anda Tidak Lulus! {{end}}`)
	if err != nil {
		panic(err)
	}
	err = t3.Execute(os.Stdout, s)
}
