package main

import (
	"fmt"
)


type User struct {
	Name string
}

func createUser(name string) *User {
	return &User{Name: name}
}


func main() {
	user := createUser("Batman")
	fmt.Println(user.Name)
	fmt.Printf("type %T", user)
}