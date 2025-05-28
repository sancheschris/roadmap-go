package main

import "fmt"

/*
	Create a map[int]User to simulate a database.

	Write a function to return all users above a certain age.

	Filter users who have verified their email.
*/

type User struct {
	ID int
	Name string
	Email string
	Age int
	Verified bool
}

func main() {

	users := map[int]User{
		10001: {1, "Christian", "chris@gmail.com", 32, true},
		10002: {2, "Edna", "edna@gmail.com", 52, true},
		10003: {3, "Gabi", "gabig@gmail.com", 26, false},
		10005: {4, "Adalberto", "adalb@gmail.com", 58, false},
	}

	printUsers(fetchAllUsersByAge(users, 29))

	printUsers(filterVerifiedUsers(users))

}

func fetchAllUsersByAge(users map[int]User, minAge int) []User {
	var result []User
	for _, user := range users {
		if user.Age > minAge {
			result = append(result, user)
		}
	}
	return result
}

func filterVerifiedUsers(users map[int]User) []User {
	var result []User
	for _, user := range users {
		if user.Verified {
			result = append(result, user)
		}
	}
	return result
}

func printUsers(users []User) {
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d, Email: %s, Verified: %v\n", user.Name, user.Age, user.Email, user.Verified)
	}
}