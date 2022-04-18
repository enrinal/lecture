package main

import (
	"fmt"
)

/*
Iteration
Sum(5) -> 5 + Sum(4) = 5 + 10 = 15
Sum(4) -> 4 + Sum(3) = 4 + 6 = 10
Sum(3) -> 3 + Sum(2) = 3 + 3 = 6
Sum(2) -> 2 + Sum(1) = 2 + 1 = 3
Sum(1) -> 1 + Sum(0) = 1 + 0 = 1
Sum(0) -> 0
*/

func Sum(number int) int {
	if number == 0 {
		return 0
	}
	return number + Sum(number-1)
}

func SumNonRecursive(number int) int {
	sum := 0
	for i := number; i > 0; i-- {
		sum += i
	}
	return sum
}

/*
Iteration
Factorial(5) -> 5 * Factorial(4) = 5 * 24 = 120
Factorial(4) -> 4 * Factorial(3) = 4 * 6 = 24
Factorial(3) -> 3 * Factorial(2) = 3 * 2 = 6
Factorial(2) -> 2 * Factorial(1) = 2 * 1 = 2
Factorial(1) -> 1 * Factorial(0) = 1 * 1 = 1
Factorial(0) -> 1
*/

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}

func main() {
	fmt.Println("Sum(5) ->", Sum(5))

	fmt.Println("Factorial(5) took", Factorial(5))
}
