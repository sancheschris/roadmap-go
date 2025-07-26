package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	println(string(body))
}