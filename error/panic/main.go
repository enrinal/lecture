package main

import "fmt"

func division(first, second int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	return first / second
}

func main() {
	fmt.Println(division(10, 2))
}
