package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://linkedin.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	defer req.Body.Close() // defer is the last code to be executed
	if err != nil {
		panic(err)
	}
	println(string(res))
	// req.Body.Close() // if you open a request, you need to close it.
}