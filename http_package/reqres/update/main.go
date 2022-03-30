package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateUsersRequest struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type UpdateUsersResponse struct {
	Name      string `json:"name"`
	Job       string `json:"job"`
	UpdatedAt string `json:"updated_at"`
}

func main() {
	var client = &http.Client{}

	req := UpdateUsersRequest{
		Name: "morpheus",
		Job:  "zion resident",
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(jsonReq)

	request, err := http.NewRequest("PUT", "https://reqres.in/api/users/2", buff)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var res UpdateUsersResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
}
