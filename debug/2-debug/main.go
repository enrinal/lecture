package main

import (
	"log"
)

func sumAll(arr []int) int {
	//fmt.Println("Processing: ", arr)
	log.Println("Processing: ", arr)

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	//fmt.Println("Result: ", sum)
	log.Println("Result: ", sum)
	return sum
}

func sum(num1, num2 int) int {
	result := num1 + num2
	return result
}

func main() {
	arr := []int{0, 1, 2, 3}
	sumAll(arr)

	arr = []int{1, 2, 3, 0}
	sumAll(arr)

	arr = []int{1, 2, 3, 4}
	sumAll(arr)

	arr = []int{4, 3, 2, 1}
	sumAll(arr)

	//println(sum(1, 2))
}
