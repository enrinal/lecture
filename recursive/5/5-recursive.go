package main

import "fmt"

func main() {
	var anonFunc func(int)
	anonFunc = func(number int) {
		if number == 0 {
			return
		} else {
			fmt.Println(number)
			anonFunc(number - 1)
		}
	}

	// call anonymous function
	anonFunc(5)
}
