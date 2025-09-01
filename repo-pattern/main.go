package main

import (
	"fmt"
	"repo-user/entity"
	"repo-user/repository"
)



func main() {
	repo := repository.InMemoryUserRepo{
		Users: make(map[int]entity.User),
	}

	user := entity.NewUser(1, "Chris", "chris@gmail.com", 32)

	user2 := &entity.User{
		ID: 2,
		Name: "Edna",
		Email: "ed@gmail.com",
		Age: 53,
	}

	repo.Save(user)
	repo.Save(user2)
	foundUser, err := repo.FindByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
	secondUser, err := repo.FindByID(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(secondUser)

	fmt.Println("----- Update user -------")

	user2.Age = 54
	repo.Update(user2)

	updatedUser, _ := repo.FindByID(2)
	fmt.Println(updatedUser)
}

