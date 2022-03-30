package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("https://postman-echo.com/post",
		url.Values{
			"name":       {"Admin"},
			"occupation": {"Engineer"},
		})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
