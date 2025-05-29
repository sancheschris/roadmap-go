package main

import (
	"fmt"
)

/*
Link orders to users and products.

Write a function that returns the total number of orders per user.

Change the status of an order.
*/

type Product struct {
	ID int
	Name string
	Price float64
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Order struct {
	OrderID string
	UserId int
	ProductIDs []int
	TotalAmount float64
	Status string
}

func main() {

	users := map[int]User{
		1: {1, "Christian", "chris@gmail.com"},
		2: {2, "Edna", "edna@gmail.com"},
	}

	products := map[int]Product{
		101: {101, "Macbook Pro", 1999.99},
		102: {102, "iPhone 16 Pro", 1099.99},
		103: {103, "iPad", 899.99},
	}

	orders := []Order{
		{"ORD001", 1, []int{101, 102}, 3099.98, "Processing"},
		{"ORD002", 2, []int{103}, 899.99, "Shipped"},
		{"ORD003", 1, []int{102}, 1099.99, "Delivered"},
	}

	userOderCount := getTotalOrdersPerUser(orders)
	fmt.Println("Total orders per user:")
	for userID, count := range userOderCount {
		fmt.Printf("User %d (%s): %d orders\n", userID, users[userID].Name, count)
	}

	// 2. Change order status
	fmt.Println("\nChanging status of order ORD002 to Delivered")
	changeOrderStatus(&orders, "ORD002", "Delivered")

	// 3. Display updated orders
	fmt.Println("\nUpdated orders:")
	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}

	fmt.Printf("\nTotal revenue from orders: $%.2f\n", getTotalRevenueOrders(orders))

	// 4. Get orders filtered by status
	filteredOrders := getOrderFilteredByStatus(orders, "Delivered")
	fmt.Println("\nOrders with status 'Delivered':")
	for _, order := range filteredOrders {
		fmt.Printf("%+v\n", order)
	}

	productNames := getProductNamesByIDs([]int{101, 102, 103}, products)
	fmt.Println("\nProduct names by IDs:")
	for _, name := range productNames {
		fmt.Println(name)
	}
	
}

func getTotalOrdersPerUser(orders []Order) map[int]int {
	result := make(map[int]int)
	for _, order := range orders {
		result[order.UserId]++
	}
	return result
}

func changeOrderStatus(orders *[]Order, orderID string, newStatus string) {
	for i, order := range *orders {
		if order.OrderID == orderID {
			(*orders)[i].Status = newStatus
			break
		}
	}
}

func getTotalRevenueOrders(orders []Order) float64 {
	var total = 0.0
	for _, order := range orders {
		total += order.TotalAmount
	}
	return total
}

func getOrderFilteredByStatus(orders []Order, status string) []Order {
	var filteredOrders []Order
	for _, order := range orders {
		if order.Status == status {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func getProductNamesByIDs(productsIDs []int, products map[int]Product) []string {
	var productNames []string
	for _, id := range productsIDs {
		if product, exists := products[id]; exists {
			productNames = append(productNames, product.Name)
		}
	}
	return productNames
}