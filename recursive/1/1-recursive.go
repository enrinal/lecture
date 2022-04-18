package main

import "fmt"

func countDown(number int) {
	fmt.Println("Countdown Starts:")
	// display the number
	fmt.Println(number)
	if number == 0 {
		return
	}
	// recursive call by decreasing number
	countDown(number - 1)
}

func main() {
	countDown(10)
}

/*
Output:
Countdown Starts: 5
5
4
3
2
1
0
*/
