package main

import "fmt"

func changeValue(x *int) {
	*x = *x + 1
}

func changeSlice(x []int) {
	x[0] = x[0] + 1
}

func mapChange(x map[string]int) {
	x["a"] = x["a"] + 1
}

func main() {
	a := 10
	changeValue(&a)
	fmt.Println(a)

	b := []int{10, 20, 30}
	changeSlice(b)
	fmt.Println(b)

	c := map[string]int{"a": 10, "b": 20}
	mapChange(c)
	fmt.Println(c)
}
