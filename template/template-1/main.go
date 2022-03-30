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
	t, err := template.New("hello").Parse("Hello, {{.}}!")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "world")
	if err != nil {
		panic(err)
	}

	// student struct
	s := Student{
		Name:  "John",
		Age:   20,
		Score: 80.5,
	}
	fmt.Println("")
	t1, err := template.New("student").Parse("Halooo {{.Name}}! Your age is {{.Age}}. Your score is {{.Score}}.")
	if err != nil {
		panic(err)
	}
	err = t1.Execute(os.Stdout, s)
	fmt.Println("")
}
