package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.Get()
	url := "https://postman-echo.com/get"
	resp, err := http.Get(fmt.Sprintf("%s?foo1=bar1&foo2=bar2", url))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(string(body))
}
