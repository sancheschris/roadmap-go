package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request initialized")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processed successfully!")
		w.Write([]byte("Request processed successfully!"))
	case <-ctx.Done():
		log.Println("Request canceled by the client")
	}
}