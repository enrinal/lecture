package main

import (
	"fmt"
	"time"
)

// compare search complexity

func LinearSearch(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	// make sorted array
	var arr []int
	for i := 0; i < 2000000; i++ {
		arr = append(arr, i)
	}

	start := time.Now()
	LinearSearch(arr, 1321122)
	end := time.Now()
	fmt.Println("Linear Search: ", end.Sub(start))

	start = time.Now()
	BinarySearch(arr, 1321122)
	end = time.Now()
	fmt.Println("Binary Search: ", end.Sub(start))
}
