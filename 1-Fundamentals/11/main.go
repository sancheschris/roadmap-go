package main

import "fmt"

type Address struct {
	City string
	State string
	Number int
}

type Client struct {
	Name string
	Age int
	Active bool
	Address
}

func main() {
	chris := Client{
		Name: "Christian",
		Age: 32,
		Active: true,
	}

	chris.City = "SP"

	fmt.Printf("The client name is %s, their age: %d, and their status is: %t and city %s", chris.Name, chris.Age, chris.Active, chris.City)
}