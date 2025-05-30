package main

import (
	"fmt"
)

type User struct {
	Name string
}

func updateName(user *User, newName string) {
	user.Name = newName
}

func main() {
	u := User{Name: "Ed"}
	updateName(&u, "Motta")
	fmt.Println(u.Name)
}