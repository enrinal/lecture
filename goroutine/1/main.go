package main

import (
	"fmt"
	"time"
)

// goroutine -> lightweight thread
// goroutine -> asyncronous
// gorutine tidak berurutan dalam pengeksekusian
func main() {
	// goroutine dengan function
	for i := 0; i < 3; i++ {
		go printHello(i)
		go printRandom(i)
	}

	// goroutine dengan anonymouse func
	go func() {
		fmt.Println("hello from anonymous func")
	}()

	// goroutine dengan closure/anonymous func
	arr := []int{1, 2, 3, 4}
	for _, val := range arr {
		go func(val int) {
			fmt.Println("value :", val)
		}(val)
	}

	time.Sleep(100 * time.Millisecond)

	// sequential
	for i := 0; i < 3; i++ {
		printHello(i)
		printRandom(i)
	}
}

func printHello(i int) {
	fmt.Println("Hello World iteration-", i)
}

func printRandom(i int) {
	fmt.Println("Hello Random iteration-", i)
}

/*
- Sequential
	--------------main---------------- |
	----printhello-------printrandom-- | core 1

- Concurrency
	-------------main-----------------  | core 1
	----printrandom-------------------  | core 2
	------printhello------------------  | core 3
	------------printhello-----------   | core 4
*/
