package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func demoPost() {
	var request = struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}{}

	client := http.Client{}

	request.Name = "morpheus"
	request.Job = "leader"

	buff, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	resp, err := client.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(buff))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response := struct {
		Name      string `json:"name"`
		Job       string `json:"job"`
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	println(response.Name)
	println(response.Job)
	println(response.ID)
	println(response.CreatedAt)
}

func demoPut() {
	var request = struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}{}

	client := http.Client{}

	request.Name = "morpheus"
	request.Job = "zion resident"

	buff, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", "https://reqres.in/api/users/2", bytes.NewBuffer(buff))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response := struct {
		Name      string `json:"name"`
		Job       string `json:"job"`
		UpdatedAt string `json:"updatedAt"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	println(response.Name)
	println(response.Job)
	println(response.UpdatedAt)

}

func demoDelete() {
	client := http.Client{}

	req, err := http.NewRequest("DELETE", "https://reqres.in/api/users/2", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code: ", resp.StatusCode)
}
func main() {
	demoPost()
	fmt.Println("")
	demoPut()
	fmt.Println("")
	demoDelete()
}
