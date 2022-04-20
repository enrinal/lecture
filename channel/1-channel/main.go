package main

import (
	"fmt"
	"time"
)

// func getProduct(a chan string) {

// }

// func getSeller(b chan string) {

// }

func main() {
	c := make(chan string)

	go func() {
		c <- "Hello"
	}() // menulis ke channel

	go func() {
		fmt.Println(<-c)
	}() // mengambil data dari channel

	x := make(chan []string)

	go func() {
		x <- []string{"Satu", "Dua", "Tiga"}
	}()

	go func() {
		fmt.Println(<-x)
		x <- []string{"Empat", "Lima", "Enam"}
	}()

	go func() {
		fmt.Println(<-x)
	}()

	//fmt.Println(<-x)

	time.Sleep(200 * time.Millisecond)
}
