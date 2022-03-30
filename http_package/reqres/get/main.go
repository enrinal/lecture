package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FetchListUsersResponse struct {
	Data       []FetchSingleUserData `json:"data"`
	Page       int                   `json:"page"`
	PerPage    int                   `json:"per_page"`
	Total      int                   `json:"total"`
	TotalPages int                   `json:"total_pages"`
}

type FetchSingleUserResponse struct {
	Data FetchSingleUserData `json:"data"`
}

type FetchSingleUserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func main() {
	baseURL := "https://reqres.in/api"
	userID := 2

	// get single
	resp, err := http.Get(baseURL + "/users/" + strconv.Itoa(userID))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("Status code is not 200")
	}

	var response FetchSingleUserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Data.ID)
	fmt.Println(response.Data.Email)
	fmt.Println(response.Data.FirstName)
	fmt.Println(response.Data.LastName)
	fmt.Println(response.Data.Avatar)

	fmt.Println("\nGet List Users")
	// get list
	page := 2
	resp, err = http.Get(baseURL + "/users?page=" + strconv.Itoa(page))
	if err != nil {
		panic(err)
	}

	var responses FetchListUsersResponse
	err = json.NewDecoder(resp.Body).Decode(&responses)
	if err != nil {
		panic(err)
	}
	for _, response := range responses.Data {
		fmt.Println(response.ID)
		fmt.Println(response.Email)
		fmt.Println(response.FirstName)
		fmt.Println(response.LastName)
		fmt.Println(response.Avatar)
	}
}
