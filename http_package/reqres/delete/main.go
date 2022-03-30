package main

import (
	"fmt"
	"net/http"
)

func main() {
	var client = &http.Client{}

	req, err := http.NewRequest("DELETE", "https://reqres.in/api/users/15", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		panic("Expected status code 204")
	}

	fmt.Println("Status code:", response.StatusCode)
}
