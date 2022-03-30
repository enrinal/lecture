package main

import (
	"html/template"
	"os"
)

type SumInput struct {
	A int
	B int
}

func hitungJumlah(a int, b int) int {
	return a + b
}

func hitungKurang(a int, b int) int {
	return a - b
}

var functionMap = template.FuncMap{
	"sum": hitungJumlah,
	"sub": hitungKurang,
}

func main() {
	const templateText = `Result: {{sum .A .B}} ResultSub: {{sub .A .B}}`
	t, err := template.New("sum").Funcs(functionMap).Parse(templateText)
	if err != nil {
		panic(err)
	}

	data := SumInput{A: 1, B: 2}
	err = t.Execute(os.Stdout, data)
}
