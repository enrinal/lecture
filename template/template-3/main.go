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
	students := []Student{
		{Name: "budi", Age: 20, Score: 100},
		{Name: "bunga", Age: 21, Score: 90},
		{Name: "john", Age: 22, Score: 80},
	}

	t, err := template.New("students").Parse(`{{range .Student}}{{.Name}} {{.Age}} {{.Score}} {{else}} No students {{end}}`)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, struct {
		Student []Student
	}{
		Student: students,
	})

	fmt.Println("")
	t2, err := template.New("students").Parse(`{{range .}}{{.Name}} {{.Age}} {{.Score}} {{else}} No students {{end}}`)
	if err != nil {
		panic(err)
	}
	err = t2.Execute(os.Stdout, students)

}
