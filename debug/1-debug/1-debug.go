package main

import "fmt"

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	var arr []int
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(findMin(arr))
}
