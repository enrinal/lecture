package main

import "fmt"

func main() {
	// a := 123
	a := 121
	output := isPalindrome(a, &a)
	fmt.Printf("Palindrom? %v", output)
}

/*
Palindrom(121)
Iteration 1
- revNum = 0 * 10 + 121 % 10 = 1
- num = 121 / 10 = 12

Iteration 2
- revNum = 1 * 10 + 12 % 10 = 12
- num = 12 / 10 = 1

Iteration 3
- revNum = 12 * 10 + 1 % 10 = 121
- num = 1 / 10 = 0
*/
func isPalindrome(x int, dup *int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	palin := isPalindrome(x/10, dup)
	*dup = *dup / 10
	lastDigit := x % 10

	if palin && *dup%10 == lastDigit {
		return true
	}
	return false
}

/*
isPalindrome(121, nil) -> isPalindrome(12, nil)
isPalindrome(12, nil) -> isPalindrome(1, nil)
isPalindrome(1, nil) -> isPalindrome(0, nil)
isPalindrome(0, nil) -> true
*/
