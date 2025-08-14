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

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("The client %s was disabled\n", c.Name)
}

func main() {
	chris := Client{
		Name: "Christian",
		Age: 32,
		Active: true,
	}

	chris.City = "SP"
	chris.Disable()

	fmt.Printf("The client name is %s, their age: %d, and their status is: %t and city %s", chris.Name, chris.Age, chris.Active, chris.City)
}