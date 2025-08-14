package main

import "fmt"

type Account struct {
	Balance int
}

func (c *Account) Available(total int) {
	c.Balance += total
	fmt.Printf("You have a total of %v dollars\n", c.Balance)
}

func main() {
	client := Account{
		Balance: 200,
	}

	client.Available(100)
	fmt.Printf("The value from struct with name %v\n", client.Balance)
}