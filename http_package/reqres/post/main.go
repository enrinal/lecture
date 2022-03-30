package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	req := CreateUserRequest{
		Name: "morpheus",
		Job:  "leader",
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(jsonReq)

	resp, err := http.Post("https://reqres.in/api/users", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response CreateUserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
