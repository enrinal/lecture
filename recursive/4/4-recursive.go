package main

import "fmt"

func tailPrintNum(n int) {
	if n > 0 {
		fmt.Println(n)

		// last statement in
		// the recursive function
		// tail recursive call
		tailPrintNum(n - 1)
	}
}

func headPrintNum(n int) {
	/*
		headPrintNum(5) -> headPrintNum(4) -> print(5)
		headPrintNum(4) -> headPrintNum(3) -> print(4)
		headPrintNum(3) -> headPrintNum(2) -> print(3)
		headPrintNum(2) -> headPrintNum(1) -> print(2)
		headPrintNum(1) -> print(1)
	*/
	if n > 0 {

		// first statement in
		// the function
		headPrintNum(n - 1)

		// printing is done at
		// returning time
		fmt.Println(n)
	}
}

func main() {
	headPrintNum(5)
	fmt.Println()
	tailPrintNum(5)
}
