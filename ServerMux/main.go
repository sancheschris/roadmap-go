package main

import "net/http"

func main() {
	mux := http.NewServeMux() // create your own multiplex to have full controll
	
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Christian!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to " + b.title))
}