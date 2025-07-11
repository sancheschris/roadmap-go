package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
		case <-time.After(5 * time.Second):
			fmt.Fprintln(w, "Hello, World!")
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println("server", err)
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}