package main

import (
	"fmt"
	"math/rand"
	"time"
)

// compare sorting algorithms

// BubbleSort Big O(n^2)
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// QuickSort Big O(n log n)
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	start := time.Now()
	var arrNum []int
	for i := 0; i < 50000; i++ {
		intn := rand.Intn(50000)
		arrNum = append(arrNum, intn)
	}
	BubbleSort(arrNum)
	end := time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	QuickSort(arrNum)
	end = time.Now()
	fmt.Println(end.Sub(start))
}
