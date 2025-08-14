package main

import "fmt"

type Client struct {
	Name string
	Age int
	Active bool
}

func main() {
	chris := Client{
		Name: "Christian",
		Age: 32,
		Active: true,
	}

	fmt.Printf("The client name is %s, their age: %d, and their status is: %t", chris.Name, chris.Age, chris.Active)
}