package main

import (
	"fmt"
)

/*
Create a slice of Products.

Write a function to print all products that are in stock.

Write a function to calculate the total value of all products.
*/

type Product struct {
	ID int
	Name string
	Price float64
	Instock bool
}

func main() {

	products := []Product{
		{ID: 1, Name: "MacBook Pro M42", Price: 1999.99, Instock: true},
		{2, "Iphone 16 Pro", 1099.99, true},
		{3, "Ipad Pro 12.9", 1299.99, false},
		{4, "Apple Watch Series 9", 399.99, true},
	}

	printProducts(products)

	fmt.Println(totalValueProducts(products))

}

func printProducts(products []Product) {
	for _, product := range products {
		fmt.Printf("Product: %v\n", product)
	}
}

func totalValueProducts(products []Product) float64 {
	var total float64 = 0.0
	for _, product := range products {
		total += product.Price
	}
	return total
}