package main

import (
	"fmt"
	"sync"
)

func getDataProduct(c chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- []string{"Product A", "Product B"}
}

func getDataSeller(c chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- []string{"Seller A", "Seller B"}
}

func getDataCust(c chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- []string{"Cust A", "Cust B"}
}

func main() {
	waitGroup := new(sync.WaitGroup)
	c := make(chan []string)

	waitGroup.Add(3)

	go func() {
		waitGroup.Wait()
		close(c)
	}()

	go getDataProduct(c, waitGroup)
	go getDataSeller(c, waitGroup)
	go getDataCust(c, waitGroup)

	// channel -> product --> seller
	for data := range c {
		fmt.Println(data)
	}
}
