package main

import (
	"sync"
)

func getDataFromUser(wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done()

	dataChan <- "Data from user"
	//fmt.Println("Data from user")
}

func getDataFromDB(wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done()

	dataChan <- "Data from DB"
	//fmt.Println("Data from DB")
}

func getDataFromCache(wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done()

	dataChan <- "Data from cache"
	//fmt.Println("Data from cache")
}

func main() {
	var wg sync.WaitGroup

	dataChan := make(chan string)

	wg.Add(3)

	go getDataFromUser(&wg, dataChan)
	go getDataFromDB(&wg, dataChan)
	go getDataFromCache(&wg, dataChan)

	go func() {
		wg.Wait()
		close(dataChan)
	}()

	for data := range dataChan {
		println(data)
	}
}
