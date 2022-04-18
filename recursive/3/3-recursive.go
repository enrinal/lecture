package main

import "fmt"

func printOne(n int) {
	if n >= 0 {
		fmt.Println("In first function:", n)
		printTwo(n - 1)
	}
}

func printTwo(n int) {
	if n >= 0 {
		fmt.Println("In second function:", n)
		printOne(n - 1)
	}
}

func main() {
	printOne(10)

	// Ini tidak akan di print
	printOne(-1)
}
